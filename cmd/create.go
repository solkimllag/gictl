package cmd

import (
	"fmt"
	"github.com/solkimllag/gictl/github"
	"log"
)

// Creates a new issue
func createIssue(newIssue *github.Issue) {
	err := github.PostIssue(userId, repo, newIssue)
	if err != nil {
		fmt.Printf("Something went wrong with creating new issue %s", newIssue.Title)
		log.Fatal(err)
	}
}
