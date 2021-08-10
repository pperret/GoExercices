package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// AddComment add a comment to an existing issue
func AddComment(username, passwd, owner, repo, issueNumber string) (*Issue, error) {
	// Create the URL
	postURL := strings.Join([]string{"https://api.github.com/repos", owner, repo, "issues", issueNumber, "comments"}, "/")

	// Get the comment using the prefered editor
	text, err := getText()
	if err != nil {
		return nil, err
	}

	// Build the issue
	var comment Comment
	comment.Body = text

	// Create the JSON request
	var request bytes.Buffer
	if err := json.NewEncoder(&request).Encode(&comment); err != nil {
		return nil, err
	}

	// Build the HTTP request
	client := &http.Client{}
	httpRequest, err := http.NewRequest("POST", postURL, &request)
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
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("Create query failed: %s", resp.Status)
	}

	// Decode the response
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil

}
