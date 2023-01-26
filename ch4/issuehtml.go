package main

import (
	"gobook/github"
	"html/template"
	"log"
	"os"
)

const issueList = template.Must(template.New("issuelist").Parse(`
	<h1>{{.TotalCount}} issues</h1>
    <table>
		<tr style='text-align: left'>
			<th>#</th>
			<th>State</th>
			<th>User</th>
			<th>Title</th>
		</tr>
		{{range .Items}}
			<tr>
				<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
				<td>{{.State}}</td>
				<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
				<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
			</tr>
		{{end}}
	</table>
`))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
	//report, err := template.New("report").
	//	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	//	Parse(templ)
	//if err != nil {
	//	log.Fatal(err)
	//}

}
