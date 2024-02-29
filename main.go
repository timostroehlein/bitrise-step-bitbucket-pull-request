package main

import (
	"fmt"
	"os"
)

func main() {
	bitbucket_access_token := os.Getenv("bitbucket_access_token")
	bitbucket_base_url := os.Getenv("bitbucket_base_url")
	bitbucket_project_key := os.Getenv("bitbucket_project_key")
	bitbucket_repository_slug := os.Getenv("bitbucket_repository_slug")
	bitbucket_pr := os.Getenv("bitbucket_pr")
	bitbucket_pr_comment_state := os.Getenv("bitbucket_pr_comment_state")
	bitbucket_pr_comment_severity := os.Getenv("bitbucket_pr_comment_severity")
	bitbucket_pr_comment_skip_if_contains := os.Getenv("bitbucket_pr_comment_skip_if_contains")
	bitbucket_pr_comment := os.Getenv("bitbucket_pr_comment")

	// Create url
	base_url := fmt.Sprintf(
		"%s/rest/api/latest/projects/%s/repos/%s",
		bitbucket_base_url,
		bitbucket_project_key,
		bitbucket_repository_slug,
	)

	// Check whether the given comment already exists
	comments, err := getComments(bitbucket_access_token, base_url, bitbucket_pr)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if bitbucket_pr_comment_skip_if_contains != "" && doesCommentExist(comments.Values, bitbucket_pr_comment_skip_if_contains) {
		fmt.Println("Skipping: already exists")
		os.Exit(0)
	}

	// Create new comment
	reqBody := AddComment{
		Severity: bitbucket_pr_comment_severity,
		State:    bitbucket_pr_comment_state,
		Text:     bitbucket_pr_comment,
	}
	err = addComment(bitbucket_access_token, base_url, bitbucket_pr, reqBody)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
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
