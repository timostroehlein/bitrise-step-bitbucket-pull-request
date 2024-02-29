package main

type CreateBranch struct {
	Name     string `json:"name"`
	CommitID int    `json:"startPoint"`
}

type CreateBranchResp struct {
	Default         bool   `json:"default"`
	DisplayID       string `json:"displayId"`
	LatestCommit    string `json:"latestCommit"`
	LatestChangeset string `json:"latestChangeset"`
	ID              string `json:"id"`
}

type PullRequest struct {
	Closed          bool          `json:"closed"`
	Version         int           `json:"version"`
	Open            bool          `json:"open"`
	Id              int           `json:"id"`
	State           string        `json:"state"`
	Locked          bool          `json:"locked"`
	HtmlDescription string        `json:"htmlDescription"`
	UpdatedDate     int           `json:"updatedDate"`
	CreatedDate     int           `json:"createdDate"`
	ClosedDate      int           `json:"closedDate"`
	Participants    []Participant `json:"participants"`
	Reviewers       []Participant `json:"reviewers"`
	Draft           bool          `json:"draft"`
	ToRef           Ref           `json:"toRef"`
	FromRef         Ref           `json:"fromRef"`
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	Links           struct{}      `json:"links"`
}

type User struct {
	DisplayName  string   `json:"displayName"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Active       bool     `json:"active"`
	EmailAddress string   `json:"emailAddress"`
	Slug         string   `json:"slug"`
	Links        struct{} `json:"links"`
	AvatarUrl    string   `json:"avatarUrl"`
}

type Repository struct {
	Name    string   `json:"name"`
	ScmId   string   `json:"scmId"`
	Slug    string   `json:"slug"`
	Project Project  `json:"project"`
	Links   struct{} `json:"links"`
}

type Project struct {
	Key       string   `json:"key"`
	AvatarUrl string   `json:"avatarUrl"`
	Avatar    string   `json:"avatar"`
	Links     struct{} `json:"links"`
}

type Participant struct {
	LastReviewedCommit string `json:"lastReviewedCommit"`
	User               User   `json:"user"`
	Role               string `json:"role"`
	Status             string `json:"status"`
	Approved           bool   `json:"approved"`
}

type Ref struct {
	Id           string     `json:"id"`
	Type         string     `json:"type"`
	DisplayId    string     `json:"displayId"`
	LatestCommit string     `json:"latestCommit"`
	Repository   Repository `json:"repository"`
}

type AddComment struct {
	Severity       string `json:"severity"`
	Version        int    `json:"version"`
	ID             int    `json:"id"`
	State          string `json:"state"`
	ThreadResolved bool   `json:"threadResolved"`
	Text           string `json:"text"`
}

type AddCommentResp struct {
	Size       int       `json:"size"`
	Limit      int       `json:"limit"`
	IsLastPage bool      `json:"isLastPage"`
	Values     []Comment `json:"values"`
}

type Comment struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	Text        string `json:"text"`
	Author      Author `json:"author"`
	CreatedDate int64  `json:"createdDate"`
	UpdatedDate int64  `json:"updatedDate"`
	Severity    string `json:"severity"`
	State       string `json:"state"`
}

type Author struct {
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	ID           int    `json:"id"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
	Slug         string `json:"slug"`
	Type         string `json:"type"`
}
