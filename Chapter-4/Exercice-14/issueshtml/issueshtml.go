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

	"GoExercices/Chapter-4/Exercice-14/github"
)

// IssuesList is the global list of Issues returned bu GitHub
type IssuesList []*github.Issue

// IssuesIndexedList is the global list of issues (indexed by their ID)
type IssuesIndexedList map[int]*github.Issue

// IssuesMilestoneList is the list of issues per milestone (identified by their ID)
type IssuesMilestoneList map[int]IssuesList

// IssuesCreatorList is the list of issues per creator (identified by their login name)
type IssuesCreatorList map[string]IssuesList

// main is the entry point of the program
func main() {

	// Check program arguments
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <owner> <repos>\n", os.Args[0])
		os.Exit(1)
	}

	// Get the global list of issues from the GitHub server
	issuesList, err := getIssuesList(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// Build the list of issues (indexed by their Number)
	issuesIndexedList := buildIndexedList(&issuesList)

	// Build the issues lists per milestone (identified by their ID)
	issuesMilestoneList := buildMilestoneList(&issuesList)

	// Build the issues lists per creator (identified by their login name)
	issuesCreatorList := buildCreatorList(&issuesList)

	// Create URL handlers
	http.Handle("/", issuesList)
	http.Handle("/milestone/", issuesMilestoneList)
	http.Handle("/creator/", issuesCreatorList)
	http.Handle("/details/", issuesIndexedList)

	// Activate web server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// getIssuesList gets the list of Issues
func getIssuesList(owner, repos string) (IssuesList, error) {

	issuesList, err := github.ListIssues(owner, repos)
	if err != nil {
		return nil, err
	}

	return IssuesList(*issuesList), nil
}

// buildIndexedList builds the lists of issues (indexed by their ID)
func buildIndexedList(list *IssuesList) IssuesIndexedList {
	indexedList := make(IssuesIndexedList)
	for _, issue := range *list {
		indexedList[issue.Number] = issue
	}
	return indexedList
}

// buildMilestoneList builds the lists of issues per milestone (identified by their ID)
func buildMilestoneList(list *IssuesList) IssuesMilestoneList {
	milestoneList := make(IssuesMilestoneList)
	for _, issue := range *list {
		if issue.Milestone != nil {
			_, ok := milestoneList[issue.Milestone.Number]
			if !ok {
				milestoneList[issue.Milestone.Number] = make(IssuesList, 0)
			}
			milestoneList[issue.Milestone.Number] = append(milestoneList[issue.Milestone.Number], issue)
		}
	}
	return milestoneList
}

// buildCreatorList builds the lists of issues per creator (identified by their login name)
func buildCreatorList(list *IssuesList) IssuesCreatorList {
	creatorList := make(IssuesCreatorList)
	for _, issue := range *list {
		_, ok := creatorList[issue.User.Login]
		if !ok {
			creatorList[issue.User.Login] = make(IssuesList, 0)
		}
		creatorList[issue.User.Login] = append(creatorList[issue.User.Login], issue)
	}
	return creatorList
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
func (issuesList IssuesList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}

	// Display the issues list
	displayIssuesList(w, &issuesList)
}

// ServeHTTP displays the list of a milestone issues (the URL contains the milestone ID)
func (milestoneList IssuesMilestoneList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}

	// Get the milestone ID from the URL
	ms := getLastElement(r.URL.Path)

	// Convert the milestone ID to an integer
	id, err := strconv.Atoi(ms)
	if err != nil {
		http.Error(w, "Invalid milestone ID", http.StatusBadRequest)
		return
	}

	// Look for the milestone in the list
	ml, ok := milestoneList[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Display the list
	displayIssuesList(w, &ml)
}

// ServeHTTP displays the creator's issues list (the URL contains the creator login name)
func (creatorList IssuesCreatorList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}

	// Get the creator from the URL
	login := getLastElement(r.URL.Path)

	// Look for the creator in the list
	cl, ok := creatorList[login]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Display the list
	displayIssuesList(w, &cl)
}

// handlerDetails displays the details of an issue (the URL contains the issue ID)
func (indexedList IssuesIndexedList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}

	// Get the issue ID from the URL
	ii := getLastElement(r.URL.Path)

	// Convert the issue ID to an integer
	id, err := strconv.Atoi(ii)
	if err != nil {
		http.Error(w, "Invalid issue ID", http.StatusBadRequest)
		return
	}

	// Look for the issue in the list
	issue, ok := indexedList[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Display the issue details
	displayIssueDetails(w, issue)
}

// displayIssuesList displays a list of issues
func displayIssuesList(out io.Writer, list *IssuesList) {
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
			<td><b>ID:</b></td>
			<td><a href='/details/{{.Number}}'>#{{.Number}}</a></td>
		</tr>
		<tr>
			<td><b>State:</b></td>
			<td>{{.State}}</td>
		</tr>
		<tr>
			<td><b>User:</b></td>
			<td><a href='/creator/{{.User.Login}}'>{{.User.Login}}</a></td>
		</tr>
		{{if .Milestone}}
		<tr>
			<td><b>Milestone:</b></td>
			<td><a href='/milestone/{{.Milestone.Number}}'>{{.Milestone.Title}}</a></td>
		</tr>
		{{end}}
		<tr>
			<td><b>Title</b>
			</td><td>{{.Title}}</td>
		</tr>
	</table>
	<p/>
	{{.Body|text2html}}
	`))

	if err := tmpl.Execute(out, issue); err != nil {
		log.Fatal(err)
	}
}
