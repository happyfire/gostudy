package main

import (
	"github.com/happyfire/gostudy/gopl/ch4/github"
	"html/template"
	"log"
	"os"
	"time"
)

//!+template
const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

//!-template

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"repo:golang/go", "is:open", "json", "decoder"}
	}

	result, err := github.SearchIssues(args)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}
