package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ListIssues queries the GitHub issue tracker.
func ListIssues(owner string, repos string) (*IssuesList, error) {

	baseURL := strings.Join([]string{IssuesURL, owner, repos, "issues"}, "/")

	num := 1
	issues := make(IssuesList, 0)

	for {
		// Build the request with parameters
		url := baseURL + "?per_page=100&page=" + strconv.Itoa(num)

		// Build the HTTP request
		client := &http.Client{}
		httpRequest, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		// Send the request to GitHub
		resp, err := client.Do(httpRequest)
		if err != nil {
			return nil, err
		}

		// We must close resp.Body on all execution paths.
		if resp.StatusCode == http.StatusNotFound {
			resp.Body.Close()
			break
		} else if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("create query failed: %s", resp.Status)
		}

		// Decode the response using the JSON decoder
		var result IssuesListResult
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		if len(result) == 0 {
			break
		}

		// Append issues to the list
		for _, issue := range result {
			issues = append(issues, &issue)
		}

		// Next page
		num++
	}
	return &issues, nil
}
