package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CloseIssue closes an issue (the issue can be already closed)
func CloseIssue(username, token, owner, repo, issueNumber string) (*Issue, error) {
	// Builds the URL
	patchURL := strings.Join([]string{URLRepos, owner, repo, "issues", issueNumber}, "/")

	// Builds the update request
	var issue Issue
	issue.State = "closed"

	// Creates the JSON request
	var requestBuffer bytes.Buffer
	if err := json.NewEncoder(&requestBuffer).Encode(&issue); err != nil {
		return nil, err
	}

	// Builds the HTTP request
	client := &http.Client{}
	httpRequest, err := http.NewRequest("PATCH", patchURL, &requestBuffer)
	if err != nil {
		return nil, err
	}

	// Sets authentication access
	httpRequest.SetBasicAuth(username, token)

	// Sets content-type
	httpRequest.Header.Add("Content-Type", "application/json")

	// Sets Accept header (recommanded by GitHub)
	httpRequest.Header.Add("Accept", "application/vnd.github.v3+json")

	// Sends the request to github
	resp, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("close query failed: %s", resp.Status)
	}

	// Decodes the github response
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
