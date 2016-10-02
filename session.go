package figo

import "net/http"

// BaseAPIEndpoint is the official API base URL
const BaseAPIEndpoint = "https://api.figo.me"

// Session is the main object for figo Connect API access
type Session struct {

	// Endpoint of figo connect API - usually set with BaseAPIEndpoint (https://api.figo.me)
	Endpoint string

	accessToken string
	hc          *http.Client
}

// NewSession creates a new figo Connect session with the given accessToken
func NewSession(accessToken string) *Session {
	return &Session{
		Endpoint:    BaseAPIEndpoint,
		accessToken: accessToken,
		hc:          &http.Client{},
	}
}
