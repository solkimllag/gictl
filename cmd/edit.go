package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/solkimllag/gictl/github"
	"log"
	"os"
	"os/exec"
)

// Updates an existing issue
func updateIssue(issueNumber *int, updatedIssue *github.Issue) {

	err := github.UpdateIssue(userId, repo, *issueNumber, updatedIssue)
	if err != nil {
		fmt.Printf("Something went wrong with updating issue: %d", issueNumber)
		log.Fatal(err)
	}
}

// Edit an issue, convenience wrapper
func edit(issueNumber *int) {
	issue, err := getIssue(issueNumber)
	if err != nil {
		fmt.Printf("Unable to fetch issue# %d\n", issueNumber)
	} else {
		updateIssue(issueNumber, editIssue(issue))
	}
}

// Opens an issue for edit in a terminal text editor
func editIssue(issue *github.Issue) *github.Issue {

	fileName := "tmp.json"
	issueJson, err := json.MarshalIndent(issue, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fileName, issueJson, 0644)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(term, "-e", editor, fileName)
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Error while editing issue.")
	} else {
		if issue.Number != 0 {
			log.Printf("Issue %d succesfully edited", issue.Number)
		}
	}
	issueJson, err = os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(issueJson, issue)
	if err != nil {
		log.Fatal(err)
	}

	return issue
}
