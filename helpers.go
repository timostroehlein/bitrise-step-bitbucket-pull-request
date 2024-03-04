package main

import "strings"

func doesCommentExist(comments []PullRequestActivity, comment_text string) (bool, Comment) {
	for _, item := range comments {
		if item.Action == "COMMENTED" && strings.Contains(item.Comment.Text, comment_text) {
			return true, item.Comment
		}
	}
	return false, Comment{}
}
