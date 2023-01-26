package main

import (
	"fmt"
	"gobook/github"
	"log"
	"os"
	"time"
)

type term int

const (
	termLastMonth term = iota
	termLastYear
	termOther
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	for term, items := range classifyByTerm(result.Items) {
		fmt.Printf("term %s:\n", getTermLabel(term))
		for _, item := range items {
			fmt.Printf("    %v #%-5d %15.15s %.55s\n", item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
	}
}

func classifyByTerm(issues []*github.Issue) map[term][]*github.Issue {
	result := make(map[term][]*github.Issue)

	for _, issue := range issues {
		term := getTerm(issue)
		result[term] = append(result[term], issue)
	}

	return result
}

func getTerm(issue *github.Issue) term {
	createdAt := issue.CreatedAt
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)
	if createdAt.After(lastMonth) {
		return termLastMonth
	}
	lastYear := now.AddDate(-1, 0, 0)
	if createdAt.After(lastYear) {
		return termLastYear
	}
	return termOther
}

func getTermLabel(term term) string {
	switch term {
	case termLastMonth:
		return "1ヶ月未満"
	case termLastYear:
		return "1年未満"
	}
	return "1年以上"
}
