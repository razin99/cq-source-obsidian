package client

type Spec struct {
	Token string `json:"token"`
	// Defaults to "cloudquery"
	UserAgent string `json:"user_agent,omitempty"`
}
