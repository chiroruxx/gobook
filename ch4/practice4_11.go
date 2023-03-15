package main

import (
	"flag"
	"fmt"
	"gobook/ch4/practice4_11"
	"os"
)

var apiKey = flag.String("key", "", "GitHub API Key")
var repo = flag.String("repo", "", "Repository to operate issues")
var title = flag.String("title", "", "Issue title to create or edit")
var number = flag.Int("number", 0, "Issue number to show or edit or close")
var shouldClose = flag.Bool("close", false, "Close issue")
var editor = flag.String("editor", "vim", "Editor to write issue body")

func main() {
	flag.Parse()

	// validation
	if *apiKey == "" {
		fmt.Fprintf(os.Stderr, "you should set api key\n")
		os.Exit(1)
	}
	if *repo == "" {
		fmt.Fprintf(os.Stderr, "you should set repository\n")
		os.Exit(1)
	}

	args := practice4_11.Args{
		ApiKey:      apiKey,
		Repo:        repo,
		Title:       title,
		Number:      number,
		ShouldClose: shouldClose,
		Editor:      editor,
	}

	switch {
	case *shouldClose:
		// close
		closed, err := practice4_11.CloseIssue(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(2)
		}

		fmt.Println(practice4_11.ConvertIssueToString(closed))
	case *title != "" && *number == 0:
		// create
		created, err := practice4_11.CreateIssue(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(2)
		}

		fmt.Println(practice4_11.ConvertIssueToString(created))
	case *title == "" && *number != 0:
		// show
		issue, err := practice4_11.GetIssue(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(2)
		}

		fmt.Println(practice4_11.ConvertIssueToString(issue))
	case *title != "" && *number != 0:
		// update
		updated, err := practice4_11.UpdateIssue(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(2)
		}

		fmt.Println(practice4_11.ConvertIssueToString(updated))
	}
}
