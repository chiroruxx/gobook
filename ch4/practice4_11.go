package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

type Issue struct {
	Number int    `json:"number,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
	State  string `json:"state,omitempty"`
}

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

	switch {
	case *shouldClose:
		// close
		closed, err := closeIssue()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(2)
		}

		fmt.Println(convertIssueToString(closed))
	case *title != "" && *number == 0:
		// create
		created, err := createIssue()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(2)
		}

		fmt.Println(convertIssueToString(created))
	case *title == "" && *number != 0:
		// show
		issue, err := getIssue()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(2)
		}

		fmt.Println(convertIssueToString(issue))
	case *title != "" && *number != 0:
		// update
		updated, err := updateIssue()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(2)
		}

		fmt.Println(convertIssueToString(updated))
	}
}

func createIssue() (*Issue, error) {
	const errorMessagePrefix = "failed to create an issue: %v\n"

	// validation
	if *title == "" {
		return nil, fmt.Errorf("you should set title\n")
	}

	// body
	body, err := getFromEditor(*editor, "")
	if err != nil {
		return nil, err
	}

	issue := Issue{Title: *title, Body: string(body)}
	requestBody, _ := json.Marshal(issue)
	resp, err := send("POST", createEndpoint(nil), bytes.NewReader(requestBody))

	if err != nil {
		return nil, fmt.Errorf(errorMessagePrefix, err)
	}
	if resp.StatusCode == http.StatusUnprocessableEntity {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf(errorMessagePrefix+"%s\n", resp.Status, body)
	}
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf(errorMessagePrefix, resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()

	return &result, nil
}

func getIssue() (*Issue, error) {
	const errorMessagePrefix = "failed to show an issue: %v\n"

	// validation
	if *number == 0 {
		return nil, fmt.Errorf("you should set number\n")
	}

	resp, err := send("GET", createEndpoint(number), nil)

	if err != nil {
		return nil, fmt.Errorf(errorMessagePrefix, err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf(errorMessagePrefix, resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}

func updateIssue() (*Issue, error) {
	const errorMessagePrefix = "failed to update an issue: %v\n"

	// validation
	if *number == 0 {
		return nil, fmt.Errorf("you should set number\n")
	}
	if *title == "" {
		return nil, fmt.Errorf("you should set title\n")
	}

	issue, err := getIssue()
	if err != nil {
		return nil, fmt.Errorf(errorMessagePrefix, err)
	}

	body, err := getFromEditor(*editor, issue.Body)
	if err != nil {
		return nil, err
	}

	issue.Title = *title
	issue.Body = string(body)
	issue.Number = 0

	requestBody, _ := json.Marshal(issue)
	resp, err := send("PATCH", createEndpoint(number), bytes.NewReader(requestBody))

	if err != nil {
		return nil, fmt.Errorf(errorMessagePrefix, err)
	}
	if resp.StatusCode == http.StatusUnprocessableEntity {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf(errorMessagePrefix+"%s\n", resp.Status, body)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf(errorMessagePrefix, resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}

func closeIssue() (*Issue, error) {
	const errorMessagePrefix = "failed to close an issue: %v\n"

	// validation
	if *number == 0 {
		return nil, fmt.Errorf("you should set number\n")
	}

	issue := Issue{State: "closed"}

	requestBody, _ := json.Marshal(issue)
	resp, err := send("PATCH", createEndpoint(number), bytes.NewReader(requestBody))

	if err != nil {
		return nil, fmt.Errorf(errorMessagePrefix, err)
	}
	if resp.StatusCode == http.StatusUnprocessableEntity {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf(errorMessagePrefix+"%s\n", resp.Status, body)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf(errorMessagePrefix, resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}

func getFromEditor(editor string, defaultContent string) ([]byte, error) {
	if editor == "" {
		editor = "vim"
	}

	tempFile := os.TempDir() + "gobook"

	if defaultContent != "" {
		if err := os.WriteFile(tempFile, []byte(defaultContent), 0666); err != nil {
			return nil, fmt.Errorf("cannot write temporary file: %v", err)
		}
	}

	cmd := exec.Command(editor, tempFile)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("cannot run editor: %v", err)
	}
	body, err := os.ReadFile(tempFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open temporary file: %v", err)
	}
	if err = os.Remove(tempFile); err != nil {
		return nil, fmt.Errorf("cannot remove temporary file: %v", err)
	}

	return body, err
}

func send(method string, url string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+*apiKey)

	client := new(http.Client)
	return client.Do(req)
}

func convertIssueToString(issue *Issue) string {
	var result string
	buffer := bytes.NewBufferString(result)

	buffer.WriteString("Number:")
	if issue.Number != 0 {
		buffer.WriteString(" " + strconv.Itoa(issue.Number))
	}
	buffer.WriteByte('\n')

	buffer.WriteString("Title: " + issue.Title + "\n")
	buffer.WriteString("State: " + issue.State + "\n")

	if issue.Body != "" {
		buffer.WriteByte('\n')
		buffer.WriteString(issue.Body)
	}

	return buffer.String()
}

func createEndpoint(number *int) string {
	endpoint := "https://api.github.com/repos/" + *repo + "/issues"
	if number != nil {
		endpoint += fmt.Sprintf("/%d", *number)
	}
	return endpoint
}
