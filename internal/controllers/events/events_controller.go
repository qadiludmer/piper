package events

import (
	"github.com/rookout/piper/internal/controllers/events/github"
)

type Controller struct {
	// GithubWebhookSecret is the secret added to this webhook via the GitHub
	// UI that identifies this call as coming from GitHub. If empty, no
	// request validation is done.
	GithubWebhookSecret []byte

	GithubRequestValidator github.RequestValidator
}
