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
	pr_title := os.Getenv("pr_title")
	pr_source_branch := os.Getenv("pr_source_branch")
	pr_target_branch := os.Getenv("pr_target_branch")
	// PR comment inputs
	create_pr_comment, _ := strconv.ParseBool(os.Getenv("create_pr_comment"))
	pr_comment_state := os.Getenv("pr_comment_state")
	pr_comment_severity := os.Getenv("pr_comment_severity")
	pr_comment_skip_if_contains := os.Getenv("pr_comment_skip_if_contains")
	pr_comment := os.Getenv("pr_comment")

	// Create url
	bitbucket_url := fmt.Sprintf(
		"%s/rest/api/latest/projects/%s/repos/%s",
		base_url,
		project_key,
		repository_slug,
	)

	if create_pr {
		req_body := PullRequest{
			Title: pr_title,
			FromRef: Ref{
				Id: "refs/heads/" + pr_source_branch,
			},
			ToRef: Ref{
				Id: "refs/heads/" + pr_target_branch,
			},
		}
		fmt.Println("Creating pull request")
		pull_request, err := createPullRequest(access_token, bitbucket_url, req_body)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if pr == "" {
			pr = strconv.Itoa(pull_request.Id)
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
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if pr_comment_skip_if_contains != "" && doesCommentExist(comments.Values, pr_comment_skip_if_contains) {
			fmt.Println("Skipping comment creation: already exists")
			os.Exit(0)
		}

		// Create new comment
		req_body := AddComment{
			Severity: pr_comment_severity,
			State:    pr_comment_state,
			Text:     pr_comment,
		}
		fmt.Println("Adding comment")
		err = addComment(access_token, bitbucket_url, pr, req_body)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
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
