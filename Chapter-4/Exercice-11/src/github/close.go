package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CloseIssue closes an issue (the issue can be already closed)
func CloseIssue(username, passwd, owner, repo, issueNumber string) (*Issue, error) {

	// Create the URL
	patchURL := strings.Join([]string{"https://api.github.com/repos", owner, repo, "issues", issueNumber}, "/")

	// Build the update request
	var issue Issue
	issue.State = "closed"

	// Create the JSON request
	var requestBuffer bytes.Buffer
	if err := json.NewEncoder(&requestBuffer).Encode(&issue); err != nil {
		return nil, err
	}

	// Build the HTTP request
	client := &http.Client{}
	httpRequest, err := http.NewRequest("PATCH", patchURL, &requestBuffer)
	if err != nil {
		return nil, err
	}

	// Set authentication access
	httpRequest.SetBasicAuth(username, passwd)

	// Set content-type
	httpRequest.Header.Add("Content-Type", "application/json")

	// Send the request to github
	resp, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Close query failed: %s", resp.Status)
	}

	// Decode the github response
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
