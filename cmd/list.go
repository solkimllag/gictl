package cmd

import (
	"fmt"
	"github.com/solkimllag/gictl/github"
	"log"
)

// Prints the list of issues from a github repo
func printIssues() {
	result, err := github.GetIssues(userId, repo)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		fmt.Printf("#%-5d %s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
