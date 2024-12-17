package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var accessToken string

func init() {
	accessToken = os.Getenv("GITHUB_TOKEN")
}

// Get http request
// userId: github user id,
// repo: github repo name,
// method: GET, POST or PATCH
func getHTTPRequest(userId string, repo string, method string, issueNumber int, issue *bytes.Buffer) (*http.Request, error) {

	reqUri := IssuesURL + "/" + userId + "/" + repo + "/issues"
	if issueNumber != 0 {
		reqUri = reqUri + "/" + strconv.Itoa(issueNumber)
	}
	var body io.Reader = nil
	if issue != nil {
		body = io.Reader(issue)
	}
	req, err := http.NewRequest(method, reqUri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.text+json")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	return req, nil
}

// Get Issues from github
func GetIssues(userId string, repo string) (*IssuesSearchResult, error) {
	req, err := getHTTPRequest(userId, repo, "GET", 0, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("GetIssues failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result.Items); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// Get a particular issue from github
func GetIssue(userId string, repo string, issueNumber int) (*Issue, error) {

	req, err := getHTTPRequest(userId, repo, "GET", issueNumber, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("GetIssue failed: %s", resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// Post a new issue to github
func PostIssue(userId string, repo string, issue *Issue) error {

	body, err := json.Marshal(issue)
	if err != nil {
		return err
	}
	req, err := getHTTPRequest(userId, repo, "POST", 0, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return fmt.Errorf("CreateIssue failed: %s", resp.Status)
	}
	return nil
}

// Update an existing github issue
func UpdateIssue(userId string, repo string, issueNumber int, issue *Issue) error {

	body, err := json.Marshal(issue)
	if err != nil {
		return err
	}
	req, err := getHTTPRequest(userId, repo, "PATCH", issueNumber, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("updateIssue failed: %s", resp.Status)
	}
	return nil
}
