// Package Issueshtml creates a HTML server for github issues
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"Chapter-4/Exercice-14/github"
)

// issuesList is the global list of issues returned from github
var issuesList *github.IssuesListResult

// issuesList is the global list of issues (indexed by their ID)
var issuesIndexedList map[int]github.Issue

// milestoneList is the list of issues per milestone (identified by their ID)
var milestoneList map[int]github.IssuesListResult

// creatorList is the list of issues per creator (identified by their login name)
var creatorList map[string]github.IssuesListResult

// main is the entry point of the program
func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <user> <password>\n", os.Args[0])
		os.Exit(1)
	}

	// Get the list from the server
	var err error
	issuesList, err = github.ListIssues(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// Build the full issues list indexed by their ID
	buildIndexedIssueList(issuesList)

	// Build the issues lists per milestone (identified by their ID)
	buildMilestoneList(issuesList)

	// Build the issues lists per creator (identified by their login name)
	buildCreatorList(issuesList)

	// Create URL handlers
	http.HandleFunc("/", handlerList)
	http.HandleFunc("/milestone/", handlerMilestone)
	http.HandleFunc("/creator/", handlerCreator)
	http.HandleFunc("/details/", handlerDetails)

	// Activate web server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// buildIndexedIssueList builds the global list of issues (identified by their ID)
func buildIndexedIssueList(list *github.IssuesListResult) {

	// Create the indexed list
	issuesIndexedList = make(map[int]github.Issue)
	for _, issue := range *list {
		issuesIndexedList[issue.Number] = issue
	}
}

// buildMilestoneList builds the lists of issues per milestone (identified by their ID)
func buildMilestoneList(list *github.IssuesListResult) {
	milestoneList = make(map[int]github.IssuesListResult)
	for _, issue := range *list {
		if issue.Milestone != nil {
			_, ok := milestoneList[issue.Milestone.Number]
			if ok == false {
				milestoneList[issue.Milestone.Number] = make(github.IssuesListResult, 0)
			}
			milestoneList[issue.Milestone.Number] = append(milestoneList[issue.Milestone.Number], issue)
		}
	}
}

// buildCreatorList builds the lists of issues per creator (identified by their login name)
func buildCreatorList(list *github.IssuesListResult) {
	creatorList = make(map[string]github.IssuesListResult)
	for _, issue := range *list {
		_, ok := creatorList[issue.User.Login]
		if ok == false {
			creatorList[issue.User.Login] = make(github.IssuesListResult, 0)
		}
		creatorList[issue.User.Login] = append(creatorList[issue.User.Login], issue)
	}
}

// getLastElement returns the last element of an URL path
func getLastElement(s string) string {
	i := strings.LastIndex(s, "/")
	if i == -1 {
		return s
	}
	return s[i+1:]
}

// handlerList displays the full list of issues
func handlerList(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	} else {
		displayIssuesList(w, issuesList)
	}
}

// handlerMilestone displays the list of a milestone issues (the URL contains the milestone ID)
func handlerMilestone(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	} else {
		// Get the milestone ID from the URL
		ms := getLastElement(r.URL.Path)

		// Convert the milestone ID to an integer
		id, err := strconv.Atoi(ms)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		// Look for the milestone in the list
		ml, ok := milestoneList[id]
		if ok == false {
			http.NotFound(w, r)
			return
		}

		// Display the list
		displayIssuesList(w, &ml)
	}
}

// handlerCreator displays the list of a creator issues (the URL contains the creator login name)
func handlerCreator(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	} else {
		// Get the creator from the URL
		login := getLastElement(r.URL.Path)

		// Look for the creator in the list
		cl, ok := creatorList[login]
		if ok == false {
			http.NotFound(w, r)
			return
		}

		// Display the list
		displayIssuesList(w, &cl)
	}
}

// handlerDetails displays the details of an issue (the URL contains the issue ID)
func handlerDetails(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	} else {
		// Get the issue ID from the URL
		ii := getLastElement(r.URL.Path)

		// Convert the issue ID to an integer
		id, err := strconv.Atoi(ii)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		// Look for the issue in the list
		issue, ok := issuesIndexedList[id]
		if ok == false {
			http.NotFound(w, r)
			return
		}

		// Display the issue
		displayIssueDetails(w, &issue)
	}
}

// displayIssuesList displays the global list of issues
func displayIssuesList(out io.Writer, list *github.IssuesListResult) {
	var tmpl = template.Must(template.New("issuelist").Parse(`
	<a href='/'>Home</a>
	<table>
	<tr style='textalign:left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Milestone</th>
	<th>Title</th>
	</tr>
	{{range $}}
	<tr>
	<td><a href='/details/{{.Number}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='/creator/{{.User.Login}}'>{{.User.Login}}</a></td>
	{{if .Milestone}}
	<td><a href='/milestone/{{.Milestone.Number}}'>{{.Milestone.Title}}</a></td>
	{{else}}
	<td> </td>
	{{end}}
	<td>{{.Title}}</td>
	</tr>
	{{end}}
	</table>
	`))
	if err := tmpl.Execute(out, list); err != nil {
		log.Fatal(err)
	}
}

// text2html replaces line breaks by <br>
func text2html(s string) template.HTML {
	return template.HTML(strings.Replace(s, "\n", "<br>", -1))
}

// displayIssueDetails displays the details of an issue
func displayIssueDetails(out io.Writer, issue *github.Issue) {
	var tmpl = template.Must(template.New("issuedetails").Funcs(template.FuncMap{"text2html": text2html}).Parse(`
	<p><a href='/'>Home</a></p>
	<p/>
	<table>
	<tr>
	<td><b>ID:</b></td><td><a href='/details/{{.Number}}'>#{{.Number}}</a></td>
	</tr>
	<tr>
	<td><b>State:</b></td><td>{{.State}}</td>
	</tr>
	<tr>
	<td><b>User:</b></td><td><a href='/creator/{{.User.Login}}'>{{.User.Login}}</a></td>
	</tr>
	{{if .Milestone}}
	<tr>
	<td><b>Milestone:</b></td><td><a href='/milestone/{{.Milestone.Number}}'>{{.Milestone.Title}}</a></td>
	</tr>
	{{end}}
	<tr>
	<td><b>Title</b></td><td>{{.Title}}</td>
	</tr>
	</table>
	<p/>
	{{.Body|text2html}}
	`))

	if err := tmpl.Execute(out, issue); err != nil {
		log.Fatal(err)
	}
}
