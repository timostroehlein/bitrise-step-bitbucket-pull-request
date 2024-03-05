package main

import (
	"strings"
)

func doesCommentExist(comments []PullRequestActivity, comment_text string) (bool, Comment) {
	for _, item := range comments {
		if item.Action == "COMMENTED" && strings.Contains(item.Comment.Text, comment_text) {
			return true, item.Comment
		}
	}
	return false, Comment{}
}

func getPullRequestReviewers(access_token string, base_url string, pr_target_branch string, pr_source_branch string) ([]Reviewer, error) {
	// Get default reviewers
	default_reviewers, err := getDefaultReviewers(access_token, base_url)
	if err != nil {
		return []Reviewer{}, err
	}

	// Find default reviewers based on the target and source branch
	var users []User
	for _, reviewer := range default_reviewers {
		if reviewer.TargetRefMatcher.Id != "refs/heads/"+pr_target_branch && reviewer.TargetRefMatcher.Id != "ANY_REF_MATCHER_ID" {
			continue
		}
		if reviewer.SourceRefMatcher.Id == "refs/heads/"+pr_source_branch || reviewer.SourceRefMatcher.Id == "ANY_REF_MATCHER_ID" {
			users = reviewer.Reviewers
			break
		}
	}
	reviewers := make([]Reviewer, len(users))
	for i, user := range users {
		reviewers[i] = Reviewer{
			User:     user,
			Status:   "UNAPPROVED",
			Role:     "REVIEWER",
			Approved: false,
		}
	}
	return reviewers, nil
}
