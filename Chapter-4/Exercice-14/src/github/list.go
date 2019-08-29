package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ListIssues queries the GitHub issue tracker.
func ListIssues(username string, passwd string) (*IssuesListResult, error) {

	num := 1

	issues := make(IssuesListResult, 0)

	for {
		// Build the request
		url := IssuesURL + "?per_page=100&page=" + strconv.Itoa(num)

		// Build the HTTP request
		client := &http.Client{}
		httpRequest, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		// Set authentication access
		if len(username) > 0 && len(passwd) > 0 {
			httpRequest.SetBasicAuth(username, passwd)
		}

		// Send the request to github
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
			return nil, fmt.Errorf("Create query failed: %s", resp.Status)
		}

		// Decode the response
		var result IssuesListResult
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		if len(result) == 0 {
			break
		}

		for _, issue := range result {
			issues = append(issues, issue)
		}

		// Next page
		num++
	}
	return &issues, nil
}
