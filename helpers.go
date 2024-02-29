package main

import "strings"

func doesCommentExist(comments []Comment, comment_text string) bool {
	for _, item := range comments {
		if strings.Contains(item.Text, comment_text) {
			return true
		}
	}
	return false
}
