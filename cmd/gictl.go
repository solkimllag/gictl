package cmd

import (
	"flag"
	"fmt"
	"github.com/solkimllag/gictl/github"
	"os"
	"os/exec"
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
	if !binaryExists(term) {
		fmt.Printf("\nTERM env is not set or %s is not an executable.\n", term)
	}
	if !binaryExists(editor) {
		fmt.Printf("\nEDITOR env is not set or %s is not an executable.\n", editor)
	}
}

func binaryExists(binary string) bool {
	_, err := exec.LookPath(binary)
	return err == nil
}

func Gictl() {
	flag.Parse()
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
		issue, err := getIssue(issueNumber)
		if err != nil {
			fmt.Printf("Unable to fetch issue# %d\n", issueNumber)
		} else {
			fmt.Printf("#%-5d %s %.55s\n",
				issue.Number, issue.User.Login, issue.Title)
		}
	case "edit":
		issue, err := getIssue(issueNumber)
		if err != nil {
			fmt.Printf("Unable to fetch issue# %d\n", issueNumber)
		} else {
			updateIssue(issueNumber, editIssue(issue))
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
