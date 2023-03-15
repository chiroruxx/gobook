package practice4_11

import (
	"bytes"
	"encoding/json"
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

type Args struct {
	ApiKey      *string
	Repo        *string
	Title       *string
	Number      *int
	ShouldClose *bool
	Editor      *string
}

func CreateIssue(args Args) (*Issue, error) {
	const errorMessagePrefix = "failed to create an issue: %v\n"

	// validation
	if *args.Title == "" {
		return nil, fmt.Errorf("you should set title\n")
	}

	// body
	body, err := getFromEditor(*args.Editor, "")
	if err != nil {
		return nil, err
	}

	issue := Issue{Title: *args.Title, Body: string(body)}
	requestBody, _ := json.Marshal(issue)
	resp, err := send("POST", createEndpoint(nil, args), bytes.NewReader(requestBody), args)

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

func GetIssue(args Args) (*Issue, error) {
	const errorMessagePrefix = "failed to show an issue: %v\n"

	// validation
	if *args.Number == 0 {
		return nil, fmt.Errorf("you should set number\n")
	}

	resp, err := send("GET", createEndpoint(args.Number, args), nil, args)

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

func UpdateIssue(args Args) (*Issue, error) {
	const errorMessagePrefix = "failed to update an issue: %v\n"

	// validation
	if *args.Number == 0 {
		return nil, fmt.Errorf("you should set number\n")
	}
	if *args.Title == "" {
		return nil, fmt.Errorf("you should set title\n")
	}

	issue, err := GetIssue(args)
	if err != nil {
		return nil, fmt.Errorf(errorMessagePrefix, err)
	}

	body, err := getFromEditor(*args.Editor, issue.Body)
	if err != nil {
		return nil, err
	}

	issue.Title = *args.Title
	issue.Body = string(body)
	issue.Number = 0

	requestBody, _ := json.Marshal(issue)
	resp, err := send("PATCH", createEndpoint(args.Number, args), bytes.NewReader(requestBody), args)

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

func CloseIssue(args Args) (*Issue, error) {
	const errorMessagePrefix = "failed to close an issue: %v\n"

	// validation
	if *args.Number == 0 {
		return nil, fmt.Errorf("you should set number\n")
	}

	issue := Issue{State: "closed"}

	requestBody, _ := json.Marshal(issue)
	resp, err := send("PATCH", createEndpoint(args.Number, args), bytes.NewReader(requestBody), args)

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

func send(method string, url string, body io.Reader, args Args) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+*args.ApiKey)

	client := new(http.Client)
	return client.Do(req)
}

func ConvertIssueToString(issue *Issue) string {
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

func createEndpoint(number *int, args Args) string {
	endpoint := "https://api.github.com/repos/" + *args.Repo + "/issues"
	if number != nil {
		endpoint += fmt.Sprintf("/%d", *number)
	}
	return endpoint
}
