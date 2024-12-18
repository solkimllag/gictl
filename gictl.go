package main

import (
	"encoding/json"
	"fmt"
	"gictl/github"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Review

var userId string
var repo string
var term string
var editor string

func init() {
	term = os.Getenv("TERM")
	editor = os.Getenv("EDITOR")
}

func main() {
	args := os.Args[1:]
	numOfArgs := len(args)
	var command string
	var issueNumber int = 0

	if numOfArgs > 0 {
		command = args[0]
	}
	if numOfArgs > 1 {
		gitRepo := strings.Split(args[1], "/")
		userId, repo = gitRepo[0], gitRepo[1]
	}
	if numOfArgs > 2 {
		issue, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Print(err)
		} else {
			issueNumber = issue
		}
	}
	if userId == "" || repo == "" {
		fmt.Println("githubuser/githubrepo is missing")
	}

	switch command {
	case "list":
		printIssues()
	case "get":
		if issueNumber != 0 {
			issue := getIssue(issueNumber)
			fmt.Printf("#%-5d %s %.55s\n",
				issue.Number, issue.User.Login, issue.Title)
		} else {
			fmt.Println("Issue number is missing")
		}
	case "edit":
		if issueNumber != 0 {
			issue := getIssue(issueNumber)
			updateIssue(issueNumber, editIssue(issue))
		} else {
			fmt.Println("Issue number is missing")
		}
	case "create":
		var issue github.Issue
		issue.Title = "Fill in title.."
		issue.Body = "..and describe it in detail"
		createIssue(editIssue(&issue))

	default:
		printHelp()
	}
}

// Print help menu
func printHelp() {
	fmt.Println()
	fmt.Print(`Usage: gictl [COMMAND] [repo_owner_id/repo_name] [issue_number]

Commands:
  list    Print list of issues.
  get     Print an issue. Issue number must be specified.
  create  Creates a new issue for the given repo.
  edit    Edit an issue. Issue number must be specified.

ENV vars
  GITHUB_TOKEN must be set for github api authentication to work. To learn about fine grained personal access tokens visit: 
		https://docs.github.com/en/rest/authentication/authenticating-to-the-rest-api?apiVersion=2022-11-28#authenticating-with-a-personal-access-token
  EDITOR must be set
  TERM   must be set

  Both edit and create commands will attempt to open a terminal text editor. For this to work, both TERM and EDITOR env vars should be set.


Examples:
  To list all github issues for github.com/solkimllag/dotfiles repo,
  run: 
  $ gictl list solkimllag/dotfiles

  To get a specific issue run:
  $ gictl get solkimllag/dotfiles 1

  To edit a specif issue run:
  $ gictl edit solkimllag/dotfiles 3

  To create a new issue run:
  $ gictl create`)

	fmt.Printf("\n\n")
}

// Fetches a particular github issue
func getIssue(issueNumber int) *github.Issue {
	issue, err := github.GetIssue(userId, repo, issueNumber)
	if err != nil {
		log.Fatal(err)
	}
	return issue
}

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

// Creates a new issue
func createIssue(newIssue *github.Issue) {
	err := github.PostIssue(userId, repo, newIssue)
	if err != nil {
		fmt.Printf("Something went wrong with creating new issue %s", newIssue.Title)
		log.Fatal(err)
	}
}

// Updates an existing issue
func updateIssue(issueNumber int, updatedIssue *github.Issue) {

	err := github.UpdateIssue(userId, repo, issueNumber, updatedIssue)
	if err != nil {
		fmt.Printf("Something went wrong with updating issue: %d", issueNumber)
		log.Fatal(err)
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
		log.Printf("Issue %d succesfully edited", issue.Number)
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
