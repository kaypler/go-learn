// go run ./issuesreport.go repo:golang/go is:open json decoder
package main

import (
	"log"
	"os"
	"text/template"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// report, err := template.New("report").
// 	Funcs(template.FuncMap{"daysAgo": daysAgo}).
// 	Parse(templ)
// if err != nil {
// 	log.Fatal(err)
// }

var report = template.Must(template.New("issuelist")).
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err I= nil {
		log.Fatal(err)
	}
}	