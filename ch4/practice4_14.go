package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type User struct {
	URL  string `json:"html_url"`
	Name string `json:"login"`
}

type Milestone struct {
	URL   string `json:"html_url"`
	Title string `json:"title"`
}

type Issue2 struct {
	Title     string    `json:"title"`
	Number    int       `json:"number"`
	URL       string    `json:"html_url"`
	User      User      `json:"user"`
	Milestone Milestone `json:"milestone"`
}

var apiKey3 = flag.String("key", "", "GitHub API Key")
var issues []Issue2

func main() {
	flag.Parse()

	if *apiKey3 == "" {
		fmt.Fprintf(os.Stderr, "you should set api key\n")
		os.Exit(1)
	}

	issues = getIssues()
	//fmt.Println(issues)
	launch()
}

func getIssues() []Issue2 {
	req, _ := http.NewRequest("GET", "https://api.github.com/repos/golang/go/issues", nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+*apiKey3)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var issues []Issue2
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return issues
}

func launch() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	const templ = `
	<h1>Issues</h1>
	<table>
		<tr>
			<th>Title</th>
			<th>User</th>
			<th>Milestone</th>
		</tr>
		{{range .}}
			<tr>
				<td><a href="{{.URL}}">{{.Title}}</a></td>
				<td><a href="{{.User.URL}}">{{.User.Name}}</a></td>
				<td><a href="{{.Milestone.URL}}">{{.Milestone.Title}}</a></td>
			</tr>
		{{end}}
	</table>
`
	t := template.Must(template.New("issues").Parse(templ))
	t.Execute(w, issues)
}
