package github

import (
	"time"
)

const IssuesURL = "https://developer.github.com/v3/search/#search-issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	Body     string    // Markdown 格式
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
