package unit_tests

import (
	"context"
	"testing"

	"github.com/machinebox/graphql"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
)

type unitTestSuite struct {
	suite.Suite
	parser         githubRepoParser
	authorProfiles map[string]*Author // only add or update, never delete
	authorIds      []string
	articles       []*RawArticle
}

func TestUnitTestSuite(t *testing.T) {
	ts := new(unitTestSuite)
	ts.parser = *NewRepoParser()
	suite.Run(t, ts)
}

func (suite *unitTestSuite) SetupSuite() {
	// parse whole repo once
	ccIds, articles, err := suite.parser.Parse(context.TODO(), "../")

	suite.Assert().Empty(err)

	suite.authorIds = ccIds
	suite.articles = articles

}

func (suite *unitTestSuite) TearDownSuite() {
}

func (s *unitTestSuite) TestAuthorMetaJson() {
	// 1. must exist within every directory under content
	// 2. every author must be a valid user on the monolith
	// 3. every author must be unique to author folder
}

func (s *unitTestSuite) TestMarkdownFilenames() {
	// 1. all markdown files must be lowercase and kebab case
}

func (s *unitTestSuite) TestMarkdownFrontmatter() {
	// 1. all markdown files should contain all required fields:
	// 		Title, Description, DatePublished, Categories, Tags, CatalogContent
	// 2. all fields should have parseable values
	// 3. categories and tags should exist in the whitelists

}

func (s *unitTestSuite) TestNonMarkdownFilesizes() {
	// all non markdown file sizes should be less than 1mb
}

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
