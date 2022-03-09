package unit_tests

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

	"github.com/adrg/frontmatter"
	"github.com/machinebox/graphql"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
)

type unitTestSuite struct {
	suite.Suite
	parser           githubRepoParser
	authorProfiles   map[string]*Author // only add or update, never delete
	authorIds        []string
	contentBasePath  string
	ccIdsLock        sync.Mutex
	tagsFileBody     string
	categoryFileBody string
}

func TestUnitTestSuite(t *testing.T) {
	ts := new(unitTestSuite)
	ts.parser = *NewRepoParser()
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

	// 1. must exist within every directory under content
	// 2. every author must be a valid user on the monolith
	// 3. every author must be unique to author folder

	contentDirPath := filepath.Join(repoPath, CONTENT_DIR_NAME)

	contentDir, err := os.ReadDir(contentDirPath)
	s.Assert().Nil(err)

	dirWg := new(sync.WaitGroup)
	s.contentBasePath, _ = filepath.Abs(contentDirPath)

	s.CheckAuthorDir(contentDir[0], dirWg)

	// for _, item := range contentDir {
	// 	// only process dirs
	// 	if item.IsDir() {
	// 		dirWg.Add(1)
	// 		s.CheckAuthorDir(item, dirWg) // from this point down, errors are logged but not returned
	// 	} else {
	// 		s.Fail("Non directory found in top level content path")
	// 	}
	// }

	// wait for all author directories to be processed
	// dirWg.Wait()

	// all final parsed author ids are unique
	// s.Assert()

}

const byteLimit int64 = 1000000

var kebabCaseRE = regexp.MustCompile("^[a-z0-9]+(-[a-z0-9]+)*$")

func (s *unitTestSuite) CheckAuthorDir(dir fs.DirEntry, dirWg *sync.WaitGroup) {
	authorDirPath := filepath.Join(s.contentBasePath, dir.Name())
	metaFilePath := filepath.Join(authorDirPath, AUTHOR_META_FILENAME)

	// verify this dir has an author_meta.json file
	authorMetaFile, err := os.ReadFile(metaFilePath)

	// There must be an author meta
	s.Assert().Nil(err)

	// verify author meta file format
	var author authorMeta
	err = json.Unmarshal(authorMetaFile, &author)

	// it must be parsable
	s.Assert().Nil(err)

	// there must be an author CcId
	s.Assert().NotNil(author.CcId)

	s.ccIdsLock.Lock()
	// save this CCID for later fetching of profile data from the monolith
	s.authorIds = append(s.authorIds, author.CcId)
	s.ccIdsLock.Unlock()

	// list the contents of this dir
	// authorDirEntries, err := os.ReadDir(authorDirPath)

	// s.Assert().NotNil(err)

	// articleWg := new(sync.WaitGroup)

	filepath.Walk(authorDirPath+"/", func(path string, info os.FileInfo, err error) error {
		// skip the root dir while walking
		if !info.IsDir() {
			fmt.Printf("%s %v bytes\n", path, info.Size())
			if strings.HasSuffix(path, ".md") {
				fmt.Print("Markdown file found...Validating frontmatter")
				// validate markdown file contents

				ss := strings.Split(path, "/")
				lastSegment := ss[len(ss)-1]
				strippedSegment := strings.TrimSuffix(lastSegment, filepath.Ext(lastSegment))
				isKebab := kebabCaseRE.MatchString(strippedSegment)
				fmt.Println(strippedSegment)
				s.Assert().True(isKebab)

				f, _ := os.OpenFile(path, os.O_RDONLY, 0655)
				defer f.Close()

				var meta articleMeta
				_, err := frontmatter.MustParse(f, &meta)
				s.Assert().Nil(err)

				// make sure input is valid
				err = validate.Struct(meta)
				s.Assert().Nil(err)

				for _, item := range meta.Categories {
					s.Assert().Contains(s.categoryFileBody, item)
				}

				for _, item := range meta.Tags {
					s.Assert().Contains(s.tagsFileBody, item)
				}

			} else {
				// assert the size is below the limit
				s.Assert().Less(info.Size(), byteLimit)
			}
		}

		return err
	})

	// // attempt to process each  file we find
	// for _, entry := range authorDirEntries {
	// 	fmt.Println(entry.Name())
	// 	// skip if not a markdown file
	// 	if !strings.HasSuffix(entry.Name(), ".md") {
	// 		continue
	//

	// 	// 	articleWg.Add(1)

	// 	// 	articlePath := filepath.Join(authorDirPath, entry.Name())
	// 	// 	articleSlug := fmt.Sprintf("%s/%s", dir.Name(), strings.Split(entry.Name(), ".md")[0])

	// 	// 	go p.processArticle(ctx, articlePath, articleSlug, author.CcId, articleWg)

	// 	// }
	// 	// articleWg.Wait()
	// }
}

// func (s *unitTestSuite) TestMarkdownFilenames() {
// 	// 	// 1. all markdown files must be lowercase and kebab case
// }

// func (s *unitTestSuite) TestMarkdownFrontmatter() {
// 	// 1. all markdown files should contain all required fields:
// 	// 		Title, Description, DatePublished, Categories, Tags, CatalogContent
// 	// 2. all fields should have parseable values
// 	// 3. categories and tags should exist in the whitelists

// }

// func (s *unitTestSuite) TestNonMarkdownFilesizes() {
// 	// all non markdown file sizes should be less than 1mb
// }

const monolithURL = "https://monolith.production-eks.codecademy.com/graphql"

type monolithQueryResponse struct {
	AuthorProfiles []monolithProfileData `json:"profiles"`
}

type monolithProfileData struct {
	CcId      string `json:"id"`
	AvatarUrl string `json:"profileImageUrl"`
	Username  string `json:"username"`
}

// FetchAndCache will retrieve all authors with the provided ccIds from the monolith
// and cache them. Authors are never removed from the cache. They are only updated
// or added.
func (s *unitTestSuite) FetchAndCache(ctx context.Context, ccIds []string) error {

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
	if err := graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		log.Ctx(ctx).Err(err)
		return err
	}

	log.Ctx(ctx).Info().Msgf("retrieved %d user profiles from monolith at %v", len(graphqlResponse.AuthorProfiles), monolithURL)
	for _, ap := range graphqlResponse.AuthorProfiles {
		s.authorProfiles[ap.CcId] = &Author{
			AvatarURL: ap.AvatarUrl,
			Username:  ap.Username,
		}
	}

	return nil
}
