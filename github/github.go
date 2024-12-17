// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import "time"

const IssuesURL = "https://api.github.com/repos"

type IssuesSearchResult struct {
	Items []*Issue
}

type Issue struct {
	Number    int       `json:"number,omitempty"`
	HTMLURL   string    `json:"html_url,omitempty"`
	Title     string    `json:"title,omitempty"`
	State     string    `json:"state,omitempty"`
	User      *User     `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Body      string    `json:"body,omitempty"` // in Markdown format
}

type User struct {
	Login   string `json:"login,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`
}

//!-
