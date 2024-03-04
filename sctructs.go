package main

type CreateBranch struct {
	Name     string `json:"name,omitempty"`
	CommitID int    `json:"startPoint,omitempty"`
}

type CreateBranchResp struct {
	Default         bool   `json:"default,omitempty"`
	DisplayID       string `json:"displayId,omitempty"`
	LatestCommit    string `json:"latestCommit,omitempty"`
	LatestChangeset string `json:"latestChangeset,omitempty"`
	ID              string `json:"id,omitempty"`
}

type PullRequest struct {
	Closed          bool     `json:"closed,omitempty"`
	Version         int      `json:"version,omitempty"`
	Open            bool     `json:"open,omitempty"`
	Id              int      `json:"id,omitempty"`
	State           string   `json:"state,omitempty"`
	Locked          bool     `json:"locked,omitempty"`
	HtmlDescription string   `json:"htmlDescription,omitempty"`
	UpdatedDate     int      `json:"updatedDate,omitempty"`
	CreatedDate     int      `json:"createdDate,omitempty"`
	ClosedDate      int      `json:"closedDate,omitempty"`
	Draft           bool     `json:"draft,omitempty"`
	ToRef           Ref      `json:"toRef,omitempty"`
	FromRef         Ref      `json:"fromRef,omitempty"`
	Title           string   `json:"title,omitempty"`
	Description     string   `json:"description,omitempty"`
	Links           struct{} `json:"links,omitempty"`
}

type Ref struct {
	Id           string `json:"id,omitempty"`
	Type         string `json:"type,omitempty"`
	DisplayId    string `json:"displayId,omitempty"`
	LatestCommit string `json:"latestCommit,omitempty"`
}

type BitbucketError struct {
	Errors []PullRequestError `json:"errors,omitempty"`
}

type PullRequestError struct {
	Context             string      `json:"context,omitempty"`
	Message             string      `json:"message,omitempty"`
	ExceptionName       string      `json:"exceptionName,omitempty"`
	ExistingPullRequest PullRequest `json:"existingPullRequest,omitempty"`
}

type AddComment struct {
	Severity       string `json:"severity,omitempty"`
	Version        int    `json:"version"`
	ID             int    `json:"id"`
	State          string `json:"state,omitempty"`
	ThreadResolved bool   `json:"threadResolved,omitempty"`
	Text           string `json:"text,omitempty"`
}

type AddCommentResp struct {
	Size       int       `json:"size,omitempty"`
	Limit      int       `json:"limit,omitempty"`
	IsLastPage bool      `json:"isLastPage,omitempty"`
	Values     []Comment `json:"values,omitempty"`
}

type Comment struct {
	ID          int    `json:"id,omitempty"`
	Version     int    `json:"version,omitempty"`
	Text        string `json:"text,omitempty"`
	Author      Author `json:"author,omitempty"`
	CreatedDate int64  `json:"createdDate,omitempty"`
	UpdatedDate int64  `json:"updatedDate,omitempty"`
	Severity    string `json:"severity,omitempty"`
	State       string `json:"state,omitempty"`
}

type Author struct {
	Name         string `json:"name,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	ID           int    `json:"id,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Active       bool   `json:"active,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Type         string `json:"type,omitempty"`
}
