package unit_tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

const (
	CONTENT_DIR_NAME     string = "content"
	AUTHOR_META_FILENAME string = "author_meta.json"
)

var (
	validate = validator.New() // needed for article metadata YAML validation
)

type authorMeta struct {
	CcId string `json:"ccID"`
}

type dateToISO struct {
	time.Time
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

type articleMeta struct {
	Title          string    `yaml:"Title" validate:"required"`
	Description    string    `yaml:"Description" validate:"required"`
	DatePublished  dateToISO `yaml:"DatePublished" validate:"required"`
	Categories     []string  `yaml:"Categories" validate:"required"`
	Tags           []string  `yaml:"Tags" validate:"required"`
	CatalogContent []string  `yaml:"CatalogContent" validate:"required"`
}

type githubRepoParser struct {
	contentBasePath string        // absolute path of the (hopefully) discovered content folder
	authorCCIds     []string      // parsed ids to fetch profile data from the monolith
	parsedArticles  []*RawArticle // parsed article data to load into the datastore
	ccIdsLock       sync.Mutex
	articlesLock    sync.Mutex
}

func NewRepoParser() *githubRepoParser {
	p := new(githubRepoParser)
	return p
}

/*
Parse will crawl and parse the article content in the provided repo and
return back data for further processing.

Returns a slice of authorCCids to be used in a subsequent batch fetch of profile data from monolith
along with a slice of pointers to rawArticles to load associated content into a DB.

Error returned means unable to even begin crawling (ie incorrect path).

All other errors are still logged but will result in skipping of
a dir / file without halting the rest of the parse process.
*/
func (p *githubRepoParser) Parse(ctx context.Context, repoPath string) ([]string, []*RawArticle, error) {

	// grab a handle on the content folder at the provided repo location
	contentDirPath := filepath.Join(repoPath, CONTENT_DIR_NAME)
	contentDir, err := os.ReadDir(contentDirPath)

	if err != nil {
		// cant open the content dir
		log.Ctx(ctx).Err(err)
		return nil, nil, err
	}

	p.contentBasePath, _ = filepath.Abs(contentDirPath)

	dirWg := new(sync.WaitGroup)

	for _, item := range contentDir {
		// only process dirs
		if item.IsDir() {
			dirWg.Add(1)
			go p.processAuthorDir(ctx, item, dirWg) // from this point down, errors are logged but not returned
		}
	}

	// wait for all author directories to be processed
	dirWg.Wait()

	msg := fmt.Sprintf("Success! Authors: %d, Articles: %d", len(p.authorCCIds), len(p.parsedArticles))
	log.Ctx(ctx).Info().Msg(msg)
	return p.authorCCIds, p.parsedArticles, nil
}

// ProcessAuthorDir processes a single author dir.
// Dir should contain a valid author_meta.json file and article markdown files.
func (p *githubRepoParser) processAuthorDir(ctx context.Context, dir fs.DirEntry, dirWg *sync.WaitGroup) {
	defer dirWg.Done()

	authorDirPath := filepath.Join(p.contentBasePath, dir.Name())
	metaFilePath := filepath.Join(authorDirPath, AUTHOR_META_FILENAME)

	// verify this dir has an author_meta.json file
	authorMetaFile, err := os.ReadFile(metaFilePath)

	if err != nil {
		// warn that author_meta was not found (skip this dir)
		log.Ctx(ctx).Warn().Msg(err.Error())
		return
	}

	// verify author meta file format
	var author authorMeta
	err = json.Unmarshal(authorMetaFile, &author)

	if err != nil {
		// warn that could not parse (skip this dir)
		log.Ctx(ctx).Warn().Msg(err.Error())
		return
	}

	if author.CcId == "" {
		// warn and skip this dir
		log.Ctx(ctx).Warn().Msg("missing ccID in author_meta.json")
		return
	}

	p.ccIdsLock.Lock()
	// save this CCID for later fetching of profile data from the monolith
	p.authorCCIds = append(p.authorCCIds, author.CcId)
	p.ccIdsLock.Unlock()

	// list the contents of this dir
	authorDirEntries, err := os.ReadDir(authorDirPath)
	if err != nil {
		// err and skip (not able to read dir contents is a big problem!)
		log.Ctx(ctx).Err(err)
		return
	}

	articleWg := new(sync.WaitGroup)

	// attempt to process each markdown file we find
	for _, entry := range authorDirEntries {
		// skip if not a markdown file
		if !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		articleWg.Add(1)

		articlePath := filepath.Join(authorDirPath, entry.Name())
		articleSlug := fmt.Sprintf("%s/%s", dir.Name(), strings.Split(entry.Name(), ".md")[0])

		go p.processArticle(ctx, articlePath, articleSlug, author.CcId, articleWg)

	}
	articleWg.Wait()
}

// ProcessArticle processes a single article within a given author dir.
// Successful result will be added to the 'ParsedArticles' slice.
func (p *githubRepoParser) processArticle(ctx context.Context, path string, slug string, ccID string, wg *sync.WaitGroup) {
	defer wg.Done()

	meta, body, err := p.parseMarkdownFile(path)
	if err != nil {
		// warn then skip this article
		log.Ctx(ctx).Warn().Msg(err.Error())
		return
	}

	// transform data into RawArticle
	raw := RawArticle{
		Title:          meta.Title,
		Slug:           slug,
		Description:    meta.Description,
		Author:         ccID,
		DatePublished:  meta.DatePublished.Time,
		Categories:     meta.Categories,
		Tags:           meta.Tags,
		Body:           string(body),
		CatalogContent: strings.Join(meta.CatalogContent, ";"),
	}

	p.articlesLock.Lock()
	// add to list
	p.parsedArticles = append(p.parsedArticles, &raw)
	p.articlesLock.Unlock()
}

// ParseMarkdownFile parses a single markdown file and returns the meta, body, and error.
// File should contain valid frontmatter.
func (p *githubRepoParser) parseMarkdownFile(articlePath string) (*articleMeta, []byte, error) {
	f, _ := os.OpenFile(articlePath, os.O_RDONLY, 0655)
	defer f.Close()

	var meta articleMeta
	body, err := frontmatter.MustParse(f, &meta)
	if err != nil {
		return nil, nil, err
	}

	// make sure input is valid
	err = validate.Struct(meta)
	if err != nil {
		return nil, nil, err
	}

	return &meta, body, err
}
