package unit_tests

import "time"

type Category struct {
	Slug string `json:"slug"`
}

type Author struct {
	AvatarURL string `json:"avatarURL"`
	Username  string `json:"username"`
}
type ArticleListItem struct {
	Title         string     `json:"title"`
	Slug          string     `json:"slug"`
	Description   string     `json:"description"`
	Authors       []Author   `json:"authors"`
	DatePublished time.Time  `json:"datePublished"`
	Categories    []Category `json:"categories"`
}
type Article struct {
	Title          string     `json:"title"`
	Slug           string     `json:"slug"`
	Description    string     `json:"description"`
	Authors        []Author   `json:"authors"`
	DatePublished  time.Time  `json:"datePublished"`
	Categories     []Category `json:"categories"`
	Tags           []string   `json:"tags"`
	Body           string     `json:"body"`
	CatalogContent []string   `json:"catalogContent"`
}

// RawArticle is an intermediary representation of an article when loading / retrieving from current DB schema.
// Must be combined with author profile data from the monolith and reformatted to 'Article' before returning out of API.
type RawArticle struct {
	Title          string    `json:"title"`
	Slug           string    `json:"slug"`
	Description    string    `json:"description"`
	Author         string    `json:"author"` // ccID (right now 1 author, keeping 'Article' authors plural for downstream flexibility)
	DatePublished  time.Time `json:"datePublished"`
	Categories     []string  `json:"categories"`
	Tags           []string  `json:"tags"`
	Body           string    `json:"body"`
	CatalogContent string    `json:"catalogContent"` // Semi-colon ; separated list of courses and paths
}

type MockArticleItemCollection map[string]ArticleListItem

type MockArticleCollection map[string]Article

type MockRawArticleCollection map[string]RawArticle
