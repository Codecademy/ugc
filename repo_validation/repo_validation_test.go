package repo_validation

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/go-playground/validator"
	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/suite"
)

const monolithURL = "https://monolith.production-eks.codecademy.com/graphql"

// size limit for non markdown files
const byteLimit int64 = 1000000

// validationr regex for markdown files
var kebabCaseRE = regexp.MustCompile("^[a-z0-9]+(-[a-z0-9]+)*$")

const (
	CONTENT_DIR_NAME     string = "content"
	AUTHOR_META_FILENAME string = "author_meta.json"
)

type monolithQueryResponse struct {
	AuthorProfiles []monolithProfileData `json:"profiles"`
}

type monolithProfileData struct {
	CcId      string `json:"id"`
	AvatarUrl string `json:"profileImageUrl"`
	Username  string `json:"username"`
}

type authorMeta struct {
	CcId string `json:"ccID"`
}

var (
	validate = validator.New() // needed for article metadata YAML validation
)

type dateToISO struct {
	time.Time
}

type articleMeta struct {
	Title          string    `yaml:"Title" validate:"required"`
	Description    string    `yaml:"Description" validate:"required"`
	DatePublished  dateToISO `yaml:"DatePublished" validate:"required"`
	Categories     []string  `yaml:"Categories" validate:"required"`
	Tags           []string  `yaml:"Tags" validate:"required"`
	CatalogContent []string  `yaml:"CatalogContent" validate:"required"`
}

// UnmarshalYAML allows conversion of date string to time.Time.
func (t *dateToISO) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return err
	}

	tt, err := time.Parse("2006-01-02", strings.TrimSpace(buf))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

type unitTestSuite struct {
	suite.Suite
	authorIds        []string
	contentBasePath  string
	ccIdsLock        sync.Mutex
	tagsFileBody     string
	categoryFileBody string
}

func TestRepoValidationSuite(t *testing.T) {
	ts := new(unitTestSuite)
	suite.Run(t, ts)
}

func (suite *unitTestSuite) SetupSuite() {
	body, err := ioutil.ReadFile("../documentation/categories.md")
	suite.Assert().Nil(err)
	suite.categoryFileBody = string(body)

	body, err = ioutil.ReadFile("../documentation/tags.md")
	suite.Assert().Nil(err)
	suite.tagsFileBody = string(body)
}

func (suite *unitTestSuite) TearDownSuite() {

}

const repoPath = "./.."

func (s *unitTestSuite) TestValidateRepo() {
	contentDirPath := filepath.Join(repoPath, CONTENT_DIR_NAME)

	contentDir, err := os.ReadDir(contentDirPath)
	s.Assert().Nil(err)

	dirWg := new(sync.WaitGroup)
	s.contentBasePath, _ = filepath.Abs(contentDirPath)

	for _, item := range contentDir {
		// only process dirs
		if item.IsDir() {
			dirWg.Add(1)
			// validate contents of directory for author
			s.validateAuthorDir(item, dirWg)
		} else {
			s.Fail("Non directory found in top level content path")
		}
	}

	// wait until all author directories are parsed so that s.authorIds is set
	dirWg.Wait()

	// assert there are no duplicate parsed author ids
	duplicateAuthorIds := hasDuplicates(s.authorIds)
	s.Assert().False(duplicateAuthorIds)

	// assert all author ids map to valid users in the production monolith
	authorData := s.fetchAuthors(s.authorIds)
	s.Assert().Equal(len(authorData.AuthorProfiles), len(s.authorIds))
}

// validateAuthorDir run validations on an author's directory by checking for a valid author_meta.json,
// valid articles, and that non articles do not exceed the size limit
func (s *unitTestSuite) validateAuthorDir(dir fs.DirEntry, dirWg *sync.WaitGroup) {
	defer dirWg.Done()

	authorDirPath := filepath.Join(s.contentBasePath, dir.Name())
	metaFilePath := filepath.Join(authorDirPath, AUTHOR_META_FILENAME)

	// verify this dir has an author_meta.json file
	authorMetaFile, err := os.ReadFile(metaFilePath)
	s.Assert().Nil(err)

	// verify author meta file format is parsable and a CcId exists
	var author authorMeta
	err = json.Unmarshal(authorMetaFile, &author)
	s.Assert().Nil(err)
	s.Assert().NotNil(author.CcId)

	// save this CCID for later fetching of profile data from the monolith
	s.ccIdsLock.Lock()
	s.authorIds = append(s.authorIds, author.CcId)
	s.ccIdsLock.Unlock()

	// wait until all articles in dir are processed
	articleWg := new(sync.WaitGroup)
	filepath.Walk(authorDirPath, func(path string, info os.FileInfo, err error) error {
		// skip the root dir while walking
		if !info.IsDir() {
			fmt.Printf("%s %v bytes\n", path, info.Size())
			if strings.HasSuffix(path, ".md") {
				articleWg.Add(1)
				fmt.Print("Markdown file found... Validating frontmatter")
				s.validateMarkdownFile(path, articleWg)
			} else {
				// assert the size is below the limit for non markdown files
				s.Assert().Less(info.Size(), byteLimit)
			}
		}

		return err
	})
}

// validateMarkdownFile runs validations on the provided markdown file path by making sure front matter is parsable and whitelisted
// and that the file name is in lowercase kebab case
func (s *unitTestSuite) validateMarkdownFile(path string, wg *sync.WaitGroup) {
	defer wg.Done()

	// validate markdown file name
	ss := strings.Split(path, "/")
	lastSegment := ss[len(ss)-1]
	strippedSegment := strings.TrimSuffix(lastSegment, filepath.Ext(lastSegment))
	isKebab := kebabCaseRE.MatchString(strippedSegment)
	fmt.Println(strippedSegment)
	s.Assert().True(isKebab)

	// validate markdown front matter
	f, _ := os.OpenFile(path, os.O_RDONLY, 0655)
	defer f.Close()

	var meta articleMeta
	_, err := frontmatter.MustParse(f, &meta)
	s.Assert().Nil(err)

	// make sure input is valid
	err = validate.Struct(meta)
	s.Assert().Nil(err)

	// ensure categories and tags are in whitelist
	for _, item := range meta.Categories {
		s.Assert().Contains(s.categoryFileBody, item)
	}

	for _, item := range meta.Tags {
		s.Assert().Contains(s.tagsFileBody, item)
	}
}

//fetchAuthors retrieves data for the given ccIds from the monolith
func (s *unitTestSuite) fetchAuthors(ccIds []string) monolithQueryResponse {
	graphqlClient := graphql.NewClient(monolithURL)
	graphqlRequest := graphql.NewRequest(`
	query ($ccIds: [String!]!){
		profiles(ccIds: $ccIds) {
			id
			profileImageUrl
			username
		}
	}`)

	graphqlRequest.Var("ccIds", ccIds)
	graphqlResponse := monolithQueryResponse{}
	err := graphqlClient.Run(context.TODO(), graphqlRequest, &graphqlResponse)
	s.Assert().Nil(err)

	return graphqlResponse
}

// hasDuplicates checks if the provided slice contains any duplicates
func hasDuplicates(items []string) bool {
	m := make(map[string]bool)
	for _, item := range items {
		if m[item] == false {
			m[item] = true
		} else {
			return true
		}
	}

	return false
}
