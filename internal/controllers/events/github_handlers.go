package events

import (
	"fmt"
	"github.com/gin-gonic/gin"
	githubSDK "github.com/google/go-github/v53/github"
	"net/http"
)

func (c *Controller) handleGithubPost(ctx *gin.Context) {
	// Validate the request against the optional webhook secret.
	payload, err := c.GithubRequestValidator.Validate(ctx.Request, c.GithubWebhookSecret)
	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	githubReqID := "X-Github-Delivery=" + ctx.Request.Header.Get("X-Github-Delivery")
	// logger := e.Logger.With("gh-request-id", githubReqID)

	event, err := githubSDK.ParseWebHook(githubSDK.WebHookType(ctx.Request), payload)
	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resp HTTPResponse

	switch _ := event.(type) {
	case *githubSDK.PushEvent:
		// TODO: (nissim)
		break
	case *githubSDK.PullRequestEvent:
		// TODO: (nissim)
		break
	default:
		resp = HTTPResponse{
			Body: fmt.Sprintf("ignoring unsupported event from github, event: %s", githubReqID),
		}
	}

	if resp.Error.Code != 0 {
		if !resp.Error.IsSilenced {
			// TODO: (nissim)
		}
		ctx.AbortWithStatusJSON(resp.Error.Code, gin.H{"error": resp.Error.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"body": resp.Body})

}
