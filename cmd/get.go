package cmd

import (
	"github.com/solkimllag/gictl/github"
	"log"
)

// Fetches a particular github issue
func getIssue(issueNumber int) *github.Issue {
	issue, err := github.GetIssue(userId, repo, issueNumber)
	if err != nil {
		log.Fatal(err)
	}
	return issue
}
