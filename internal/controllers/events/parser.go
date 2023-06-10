package events

import "github.com/google/go-github/v53/github"

type PullRequestWebhook struct {
}

type PushWebhook struct {
}

// Parser parses webhook events from different VCS hosts into our models
type Parser interface {
	// ParseGithubPullEvent parses GitHub pull request events.
	// pull is the parsed pull request.
	ParseGithubPullEvent(pullEvent *github.PullRequestEvent) (pr PullRequestWebhook, err error)

	// ParseGithubPushEvent parses GitHub push events.
	ParseGithubPushEvent(pushEvent *github.PushEvent) (push PushWebhook, err error)
}
