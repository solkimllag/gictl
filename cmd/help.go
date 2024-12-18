package cmd

import "fmt"

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
