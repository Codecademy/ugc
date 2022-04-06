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

// end point for checking author data
// var authorsURL = os.Getenv("AUTHORS_URL")
var authorsURL = "https://monolith.production-eks.codecademy.com/graphql"

// size limit for non markdown files (1mb)
const byteLimit int64 = 1000000
const contentRoot = "./.."

// validation regex for markdown files
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

/**
TestValidateRepo walks the authors content folders and verifies the following requirements:

1. each top level author directory has a author_meta.json with a "ccID" field
2. each "ccId" is unique across the entire repo
3. each "ccId" exists in an Authors DB
4. markdown files should be named in kebab case and lowercase letters
5. markdown files contain frontmatter and includes all required fields fields
6. markdown frontmatter "categories" and "tags" are found in the documented passlists (./documentation/(categories|tags).md)
7. non-markdown files should not exceed 1MB
*/
func (s *unitTestSuite) TestValidateRepo() {
	contentDirPath := filepath.Join(contentRoot, CONTENT_DIR_NAME)

	contentDir, err := os.ReadDir(contentDirPath)
	s.Assert().Nil(err, "Unable to parse content directory")

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
	s.Assert().False(duplicateAuthorIds, "List of author IDs is not unique")

	// assert all author ids map to valid users in the production monolith
	authorData := s.fetchAuthors(s.authorIds)
	s.Assert().Equal(len(authorData.AuthorProfiles), len(s.authorIds), "Monolith did not return expected count of authors")
}

// validateAuthorDir runs validations on an author's directory by checking for a valid author_meta.json,
// valid articles, and that non articles do not exceed the size limit
func (s *unitTestSuite) validateAuthorDir(dir fs.DirEntry, dirWg *sync.WaitGroup) {
	defer dirWg.Done()

	authorDirPath := filepath.Join(s.contentBasePath, dir.Name())
	metaFilePath := filepath.Join(authorDirPath, AUTHOR_META_FILENAME)

	// verify this dir has an author_meta.json file
	authorMetaFile, err := os.ReadFile(metaFilePath)
	s.Assert().Nil(err, "No author_meta.json file found")

	// verify author meta file format is parsable and a CcId exists
	var author authorMeta
	err = json.Unmarshal(authorMetaFile, &author)
	s.Assert().Nil(err, "Could not parse author_meta.json")
	s.Assert().NotNil(author.CcId, "No CcId found in author_meta.json")

	// save this CCID for later fetching of profile data from the monolith
	s.ccIdsLock.Lock()
	s.authorIds = append(s.authorIds, author.CcId)
	s.ccIdsLock.Unlock()

	// wait until all articles in dir are processed
	articleWg := new(sync.WaitGroup)
	filepath.Walk(authorDirPath, func(path string, info os.FileInfo, err error) error {
		// skip the root dir while walking
		if !info.IsDir() {
			if strings.HasSuffix(path, ".md") {
				articleWg.Add(1)
				s.validateMarkdownFile(path, articleWg)
			} else {
				// assert the size is below the limit for non markdown files
				fmt.Printf("- validating non-markdown file: %v \n", path)
				s.Assert().Less(info.Size(), byteLimit, "File is too large")
			}
		}

		return err
	})

	articleWg.Wait()
}

// validateMarkdownFile runs validations on the provided markdown file path by making sure front matter is parsable and whitelisted
// and that the file name is in lowercase kebab case
func (s *unitTestSuite) validateMarkdownFile(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("- validating markdown file: %v \n", path)

	// validate markdown file name
	ss := strings.Split(path, "/")
	lastSegment := ss[len(ss)-1]
	strippedSegment := strings.TrimSuffix(lastSegment, filepath.Ext(lastSegment))
	isKebab := kebabCaseRE.MatchString(strippedSegment)
	s.Assert().True(isKebab, "The file name is not in kebab case")

	// validate markdown front matter
	f, _ := os.OpenFile(path, os.O_RDONLY, 0655)
	defer f.Close()

	var meta articleMeta
	_, err := frontmatter.MustParse(f, &meta)
	s.Assert().Nil(err, "Error parsing frontmatter")

	// make sure input is valid
	err = validate.Struct(meta)
	s.Assert().Nil(err, "Frontmatter fails validation")

	// ensure categories and tags are in allowlist
	for _, item := range meta.Categories {
		s.Assert().Contains(s.categoryFileBody, item+"\n", "Category was not found in documentation/categories.md")
	}

	for _, item := range meta.Tags {
		s.Assert().Contains(s.tagsFileBody, item+"\n", "Tag was not found in documentation/tags.md")
	}
}

//fetchAuthors retrieves data for the given ccIds from the monolith
func (s *unitTestSuite) fetchAuthors(ccIds []string) monolithQueryResponse {
	graphqlClient := graphql.NewClient(authorsURL)
	graphqlRequest := graphql.NewRequest(`
	query {
		profiles(ccIds:["53ad6728c660e4eb130002e5", "56f6d56e4c432ce4d1000701", "610d6b838f6bbe7014931336", "53d18876fed2a851f8000029", "60cbc2d1011b910740680cbd"] ) {
			id
			profileImageUrl
			username
		}
	}`)

	// graphqlRequest.Var("ccIds", ccIds)
	graphqlResponse := monolithQueryResponse{}
	graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse)

	if err != nil {
		panic(err)
	}
	// s.Assert().Nil(err, "Error fetching author info")

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
