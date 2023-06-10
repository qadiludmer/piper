package events

type HTTPResponse struct {
	Error HTTPResponseError `json:"error,omitempty"`
	Body  string            `json:"body,omitempty"`
}

type HTTPResponseError struct {
	Code       int   `json:"code,omitempty"`
	Error      error `json:"error,omitempty"`
	IsSilenced bool  `json:"isSilenced,omitempty"`
}
