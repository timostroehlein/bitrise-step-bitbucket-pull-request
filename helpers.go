package main

import (
	"fmt"
	"path/filepath"
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
	var exact_users []User
	var any_users []User
	for _, reviewer := range default_reviewers {
		source_branch_match, err := filepath.Match(reviewer.SourceRefMatcher.Id, pr_source_branch)
		target_branch_match, err := filepath.Match(reviewer.TargetRefMatcher.Id, "refs/heads/"+pr_target_branch)
		if err != nil {
			fmt.Println("Could not match default reviewer pattern", err)
			continue
		}
		if source_branch_match && target_branch_match {
			exact_users = reviewer.Reviewers
			break
		} else if (source_branch_match || reviewer.SourceRefMatcher.Id == "ANY_REF_MATCHER_ID") && (target_branch_match && reviewer.TargetRefMatcher.Id == "ANY_REF_MATCHER_ID") {
			any_users = reviewer.Reviewers
			break
		}
	}
	var users []User
	if len(exact_users) > 0 {
		users = exact_users
	} else {
		users = any_users
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
