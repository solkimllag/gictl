package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
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

	issueNumber := flag.Int("i", 0, "Specify issue number.")
	flag.Func("gr", "Specify github repo. For example: solkimllag/gictl", func(gr string) error {
		gitRepo := strings.Split(gr, "/")
		if len(gitRepo) < 2 {
			return errors.New("Unable to parse repo name.")
		}
		userId, repo = gitRepo[0], gitRepo[1]
		if userId == "" || repo == "" {
			return errors.New("Githubuser/githubrepo is missing")
		}
		return nil
	})
	flag.Usage = func() {
		fmt.Println()
		fmt.Print(`Usage: gictl [-i n] [-gr repo_owner_id/repo_name] [command]
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
  $ gictl -gr solkimllag/dotfiles list

  To get a specific issue run:
  $ gictl -gr solkimllag/dotfiles -i 1123 get

  To edit a specif issue run:
  $ gictl -gr solkimllag/dotfiles -i 1234 edit

  To create a new issue run:
  $ gictl create

Flags:`)

		fmt.Println()
		flag.PrintDefaults()
		fmt.Println()
	}

	flag.Parse()
	switch flag.Arg(0) {
	case "list":
		printIssues()
	case "get":
		printIssue(issueNumber)
	case "edit":
		edit(issueNumber)
	case "create":
		create()
	default:
		flag.Usage()
	}
}
