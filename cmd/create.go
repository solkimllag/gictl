package cmd

import (
	"fmt"
	"github.com/solkimllag/gictl/github"
	"log"
)

// Convenience wrapper for createIssue
func create() {
	var issue github.Issue
	issue.Title = "Fill in title.."
	issue.Body = "..and describe it in detail"
	createIssue(editIssue(&issue))
}

// Creates a new issue
func createIssue(newIssue *github.Issue) {
	err := github.PostIssue(userId, repo, newIssue)
	if err != nil {
		fmt.Printf("Something went wrong with creating new issue %s", newIssue.Title)
		log.Fatal(err)
	}
}
