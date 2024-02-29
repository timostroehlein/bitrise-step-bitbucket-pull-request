package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func createBranch(access_token string, base_url string, branch CreateBranch) (CreateBranchResp, error) {
	create_branch_url := base_url + "/branches"

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(branch)
	if err != nil {
		return CreateBranchResp{}, err
	}

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", create_branch_url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return CreateBranchResp{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+access_token)

	// Send request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return CreateBranchResp{}, err
	}
	defer resp.Body.Close()

	// Handle the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CreateBranchResp{}, err
	}
	fmt.Println("Response Status:", resp.Status)
	if resp.StatusCode >= 400 {
		return CreateBranchResp{}, errors.New(string(respBody))
	}
	var createBranchResp CreateBranchResp
	err = json.Unmarshal(respBody, &createBranchResp)
	if err != nil {
		return CreateBranchResp{}, err
	}
	return createBranchResp, nil
}

func createPullRequest(access_token string, base_url string, pull_request PullRequest) (PullRequest, error) {
	create_pull_request_url := base_url + "/pull-requests"

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(pull_request)
	if err != nil {
		return PullRequest{}, err
	}

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", create_pull_request_url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return PullRequest{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+access_token)

	// Send request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return PullRequest{}, err
	}
	defer resp.Body.Close()

	// Handle the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return PullRequest{}, err
	}
	fmt.Println("Response Status:", resp.Status)
	if resp.StatusCode >= 400 {
		return PullRequest{}, errors.New(string(respBody))
	}
	var pullRequest PullRequest
	err = json.Unmarshal(respBody, &pullRequest)
	if err != nil {
		return PullRequest{}, err
	}
	return pullRequest, nil
}

func getComments(access_token string, base_url string, pull_request string) (AddCommentResp, error) {
	get_comments_url := fmt.Sprintf(
		"%s/pull-requests/%s/blocker-comments",
		base_url,
		pull_request,
	)

	params := url.Values{}
	params.Set("limit", "100")
	combined_url, err := url.Parse(get_comments_url)
	if err != nil {
		return AddCommentResp{}, err
	}
	combined_url.RawQuery = params.Encode()

	// Create the HTTP GET request
	req, err := http.NewRequest("GET", combined_url.String(), nil)
	if err != nil {
		return AddCommentResp{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+access_token)

	// Send request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return AddCommentResp{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return AddCommentResp{}, err
	}

	// Handle the response
	fmt.Println("Response Status:", resp.Status)
	if resp.StatusCode >= 400 {
		return AddCommentResp{}, errors.New(string(respBody))
	}
	var addCommentResp AddCommentResp
	err = json.Unmarshal(respBody, &addCommentResp)
	if err != nil {
		return AddCommentResp{}, err
	}

	return addCommentResp, nil
}

func addComment(access_token string, base_url string, pull_request string, comment AddComment) error {
	add_comment_url := fmt.Sprintf(
		"%s/pull-requests/%s/blocker-comments",
		base_url,
		pull_request,
	)

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(comment)
	if err != nil {
		return err
	}

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", add_comment_url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+access_token)

	// Send request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("Response Status:", resp.Status)
	if resp.StatusCode >= 400 {
		return errors.New(string(respBody))
	}
	return nil
}
