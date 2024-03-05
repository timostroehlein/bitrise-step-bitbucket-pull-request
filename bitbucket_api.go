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

func httpRequest(method string, access_token string, full_url string, request_body interface{}) ([]byte, int, error) {
	// Convert the request body to JSON
	json_body, err := json.Marshal(request_body)
	if err != nil {
		return []byte{}, 1, err
	}

	// Create the HTTP request
	req, err := http.NewRequest(method, full_url, bytes.NewBuffer(json_body))
	if err != nil {
		return []byte{}, 1, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+access_token)

	// Send request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, 1, err
	}
	defer resp.Body.Close()

	// Handle the response
	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, 1, err
	}
	fmt.Println("Response Status:", resp.Status)
	if resp.StatusCode >= 400 {
		return []byte{}, resp.StatusCode, errors.New(string(resp_body))
	}
	return resp_body, resp.StatusCode, nil
}

func createPullRequest(access_token string, base_url string, pull_request PullRequest) (PullRequest, int, error) {
	create_pull_request_url := base_url + "/pull-requests"
	resp_body, result_code, err := httpRequest("POST", access_token, create_pull_request_url, pull_request)
	if err != nil {
		var bitbucket_error BitbucketError
		err := json.Unmarshal([]byte(err.Error()), &bitbucket_error)
		if err != nil {
			return PullRequest{}, result_code, err
		}
		if len(bitbucket_error.Errors) > 0 {
			if bitbucket_error.Errors[0].ExistingPullRequest.Id != 0 {
				fmt.Println("Skipping PR creation: already exists")
				return bitbucket_error.Errors[0].ExistingPullRequest, result_code, nil
			}
			return PullRequest{}, result_code, errors.New(bitbucket_error.Errors[0].Message)
		}
		return PullRequest{}, result_code, err
	}

	var pull_request_resp PullRequest
	err = json.Unmarshal(resp_body, &pull_request_resp)
	if err != nil {
		return PullRequest{}, result_code, err
	}
	return pull_request_resp, result_code, nil
}

func updatePullRequest(access_token string, base_url string, pull_request_id string, pull_request PullRequest) error {
	update_pull_request_url := fmt.Sprintf(
		"%s/pull-requests/%s",
		base_url,
		pull_request_id,
	)
	_, _, err := httpRequest("PUT", access_token, update_pull_request_url, pull_request)
	if err != nil {
		return err
	}
	return nil
}

func getComments(access_token string, base_url string, pull_request string) (AddCommentResp, error) {
	get_comments_url := fmt.Sprintf(
		"%s/pull-requests/%s/activities",
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

	resp_body, _, err := httpRequest("GET", access_token, combined_url.String(), nil)
	if err != nil {
		return AddCommentResp{}, err
	}

	var add_comment_resp AddCommentResp
	err = json.Unmarshal(resp_body, &add_comment_resp)
	if err != nil {
		return AddCommentResp{}, err
	}
	return add_comment_resp, nil
}

func addComment(access_token string, base_url string, pull_request string, comment AddComment) error {
	add_comment_url := fmt.Sprintf(
		"%s/pull-requests/%s/blocker-comments",
		base_url,
		pull_request,
	)
	_, _, err := httpRequest("POST", access_token, add_comment_url, comment)
	if err != nil {
		return err
	}
	return nil
}

func updateComment(access_token string, base_url string, pull_request string, comment_id string, comment AddComment) error {
	update_comment_url := fmt.Sprintf(
		"%s/pull-requests/%s/comments/%s",
		base_url,
		pull_request,
		comment_id,
	)
	_, _, err := httpRequest("PUT", access_token, update_comment_url, comment)
	if err != nil {
		return err
	}
	return nil
}

func getDefaultReviewers(access_token string, base_url string) ([]DefaultReviewersResp, error) {
	default_reviewers_url := base_url + "/conditions"
	resp_body, _, err := httpRequest("GET", access_token, default_reviewers_url, nil)
	if err != nil {
		return []DefaultReviewersResp{}, err
	}

	var default_reviewers_resp []DefaultReviewersResp
	err = json.Unmarshal(resp_body, &default_reviewers_resp)
	if err != nil {
		return []DefaultReviewersResp{}, err
	}
	return default_reviewers_resp, nil
}
