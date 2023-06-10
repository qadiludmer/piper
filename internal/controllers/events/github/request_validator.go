package github

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	githubSDK "github.com/google/go-github/v53/github"
)

// RequestValidator handles checking if GitHub requests are signed
// properly by the secret.
type RequestValidator interface {
	// Validate returns the JSON payload of the request.
	// If secret is not empty, it checks that the request was signed
	// by secret and returns an error if it was not.
	// If secret is empty, it does not check if the request was signed.
	Validate(r *http.Request, secret []byte) ([]byte, error)
}

// DefaultRequestValidator handles checking if GitHub requests are signed
// properly by the secret.
type DefaultRequestValidator struct{}

// Validate returns the JSON payload of the request.
// If secret is not empty, it checks that the request was signed
// by secret and returns an error if it was not.
// If secret is empty, it does not check if the request was signed.
func (d *DefaultRequestValidator) Validate(r *http.Request, secret []byte) ([]byte, error) {
	if len(secret) != 0 {
		return d.validateAgainstSecret(r, secret)
	}
	return d.validateWithoutSecret(r)
}

// validateAgainstSecret validates the request against the provided secret.
// It returns the JSON payload of the request if the validation is successful.
func (d *DefaultRequestValidator) validateAgainstSecret(r *http.Request, secret []byte) ([]byte, error) {
	// ValidatePayload verifies if the request was signed by the provided secret.
	payload, err := githubSDK.ValidatePayload(r, secret)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// validateWithoutSecret validates the request without a secret.
// It returns the JSON payload of the request if the validation is successful.
func (d *DefaultRequestValidator) validateWithoutSecret(r *http.Request) ([]byte, error) {
	switch ct := r.Header.Get("Content-Type"); ct {
	case "application/json":
		// Read the JSON payload from the request body.
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, fmt.Errorf("could not read body: %s", err)
		}
		return payload, nil
	case "application/x-www-form-urlencoded":
		// GitHub stores the JSON payload as a form value.
		payloadForm := r.FormValue("payload")
		if payloadForm == "" {
			return nil, errors.New("webhook request did not contain the expected 'payload' form value")
		}
		return []byte(payloadForm), nil
	default:
		return nil, fmt.Errorf("webhook request has unsupported Content-Type %q", ct)
	}
}
