package git

import (
	"net/http"

	"github.com/google/go-github/v52/github"
)

type CommitFile struct {
	Path    *string `json:"path"`
	Content *string `json:"content"`
}

type WebhookPayload struct {
	Event            string `json:"event"`
	Repo             string `json:"repoName"`
	Branch           string `json:"branch"`
	Commit           string `json:"commit"`
	User             string `json:"user"`
	UserEmail        string `json:"user_email"`
	PullRequestURL   string `json:"pull_request_url"`
	PullRequestTitle string `json:"pull_request_title"`
	DestBranch       string `json:"dest_branch"`
}

type Client interface {
	ListFiles(repo string, branch string, path string) ([]string, error)
	GetFile(repo string, branch string, path string) (*CommitFile, error)
	SetWebhook() error
	UnsetWebhook() error
	GetWebhook(hookID int64) (*github.Hook, error)
	HandlePayload(request *http.Request, secret []byte) (*WebhookPayload, error)
}
