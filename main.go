package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// General inputs
	access_token := os.Getenv("access_token")
	base_url := os.Getenv("base_url")
	project_key := os.Getenv("project_key")
	repository_slug := os.Getenv("repository_slug")
	pr := os.Getenv("pr")
	// PR inputs
	create_pr, _ := strconv.ParseBool(os.Getenv("create_pr"))
	update_pr, _ := strconv.ParseBool(os.Getenv("update_pr"))
	pr_title := os.Getenv("pr_title")
	pr_description := os.Getenv("pr_description")
	pr_source_branch := os.Getenv("pr_source_branch")
	pr_target_branch := os.Getenv("pr_target_branch")
	// PR comment inputs
	create_pr_comment, _ := strconv.ParseBool(os.Getenv("create_pr_comment"))
	pr_comment_state := os.Getenv("pr_comment_state")
	pr_comment_severity := os.Getenv("pr_comment_severity")
	pr_comment_match_action := os.Getenv("pr_comment_match_action")
	pr_comment_match_string := os.Getenv("pr_comment_match_string")
	pr_comment := os.Getenv("pr_comment")

	if access_token == "" {
		fmt.Println("Missing access token")
		os.Exit(1)
	}
	if base_url == "" {
		fmt.Println("Missing base url")
		os.Exit(1)
	}
	if project_key == "" {
		fmt.Println("Missing project key")
		os.Exit(1)
	}

	// Create url
	bitbucket_url := fmt.Sprintf(
		"%s/rest/api/latest/projects/%s/repos/%s",
		base_url,
		project_key,
		repository_slug,
	)
	bitbucket_default_reviewers_url := fmt.Sprintf(
		"%s/rest/default-reviewers/latest/projects/%s/repos/%s",
		base_url,
		project_key,
		repository_slug,
	)

	new_pr_created := false
	var pr_version int
	if create_pr {
		// Get default reviewers
		fmt.Println("Requesting default reviewers")
		default_reviewers, err := getPullRequestReviewers(access_token, bitbucket_default_reviewers_url, pr_target_branch, pr_source_branch)
		if err != nil {
			fmt.Println("Get default reviewers error:", err)
		}

		// Create pull request
		req_body := PullRequest{
			Title:       pr_title,
			Description: pr_description,
			FromRef: Ref{
				Id: "refs/heads/" + pr_source_branch,
			},
			ToRef: Ref{
				Id: "refs/heads/" + pr_target_branch,
			},
			Reviewers: default_reviewers,
		}
		fmt.Println("Creating pull request")
		pull_request, result_code, err := createPullRequest(access_token, bitbucket_url, req_body)
		if err != nil {
			fmt.Println("Create pull request error:", err)
			os.Exit(1)
		}
		if result_code != 409 {
			new_pr_created = true
		}
		if pr == "" {
			pr = strconv.Itoa(pull_request.Id)
		}
		pr_version = *pull_request.Version
	}

	// Only update pr if it already exists
	if update_pr && !new_pr_created {
		if pr == "" {
			fmt.Println("Cannot update PR, missing id, make sure the given PR exists")
			os.Exit(1)
		}
		// Get default reviewers
		fmt.Println("Requesting default reviewers")
		default_reviewers, err := getPullRequestReviewers(access_token, bitbucket_default_reviewers_url, pr_target_branch, pr_source_branch)
		if err != nil {
			fmt.Println("Get default reviewers error:", err)
			err = nil
		}

		// Update pull request
		req_body := PullRequest{
			Version:     &pr_version,
			Title:       pr_title,
			Description: pr_description,
			FromRef: Ref{
				Id: "refs/heads/" + pr_source_branch,
			},
			ToRef: Ref{
				Id: "refs/heads/" + pr_target_branch,
			},
			Reviewers: default_reviewers,
		}
		err = updatePullRequest(access_token, bitbucket_url, pr, req_body)
		if err != nil {
			fmt.Println("Update pull request error:", err)
			os.Exit(1)
		}
	}

	if create_pr_comment {
		if pr == "" {
			fmt.Println("Cannot create PR comment, missing id, make sure the given PR exists")
			os.Exit(1)
		}

		// Check whether the given comment already exists
		fmt.Println("Requesting comments")
		comments, err := getComments(access_token, bitbucket_url, pr)
		if err != nil {
			fmt.Println("Get comments error:", err)
			os.Exit(1)
		}

		comment_exists, existing_comment := doesCommentExist(comments.Values, pr_comment_match_string)
		if pr_comment_match_string != "" && comment_exists {
			if pr_comment_match_action == "SKIP" {
				fmt.Println("Skipping comment creation: already exists")
				os.Exit(0)
			}

			// Update existing comment
			req_body := AddComment{
				Version:  existing_comment.Version,
				Severity: pr_comment_severity,
				State:    pr_comment_state,
				Text:     pr_comment,
			}
			fmt.Println("Updating existing comment")
			err = updateComment(access_token, bitbucket_url, pr, strconv.Itoa(existing_comment.ID), req_body)
			if err != nil {
				fmt.Println("Update comment error:", err)
				os.Exit(1)
			}
		} else {
			// Create new comment
			req_body := AddComment{
				Severity: pr_comment_severity,
				State:    pr_comment_state,
				Text:     pr_comment,
			}
			fmt.Println("Adding comment")
			err = addComment(access_token, bitbucket_url, pr, req_body)
			if err != nil {
				fmt.Println("Add comment error:", err)
				os.Exit(1)
			}
		}
	}

	// --- Step Outputs: Export Environment Variables for other Steps:
	// You can export Environment Variables for other Steps with
	//  envman, which is automatically installed by `bitrise setup`.
	// A very simple example:
	// cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", "EXAMPLE_STEP_OUTPUT", "--value", "the value you want to share").CombinedOutput()
	// if err != nil {
	// 	fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
	// 	os.Exit(1)
	// }
	// You can find more usage examples on envman's GitHub page
	//  at: https://github.com/bitrise-io/envman

	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
