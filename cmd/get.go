package cmd

import (
	"fmt"
	"github.com/solkimllag/gictl/github"
)

// Prints a particular github issue
func printIssue(issueNumber *int) {
	issue, err := getIssue(issueNumber)
	if err != nil {
		fmt.Printf("Unable to fetch issue# %d\n", issueNumber)
	} else {
		fmt.Printf("#%-5d %s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}
}

// Fetches a particular github issue
func getIssue(issueNumber *int) (*github.Issue, error) {
	issue, err := github.GetIssue(userId, repo, *issueNumber)
	return issue, err
}
