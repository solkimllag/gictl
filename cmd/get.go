package cmd

import (
	"github.com/solkimllag/gictl/github"
)

// Fetches a particular github issue
func getIssue(issueNumber int) (*github.Issue, error) {
	issue, err := github.GetIssue(userId, repo, issueNumber)
	return issue, err
}
