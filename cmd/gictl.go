package cmd

import (
	"fmt"
	"github.com/solkimllag/gictl/github"
	"os"
	"strconv"
	"strings"
)

var userId string
var repo string
var term string
var editor string

func init() {
	term = os.Getenv("TERM")
	editor = os.Getenv("EDITOR")
}

func Gictl() {
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
