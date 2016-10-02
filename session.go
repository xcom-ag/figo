package figo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

// GetUser retrieves the current user
func (s *Session) GetUser() (*User, error) {
	req, err := http.NewRequest("GET", s.Endpoint+"/rest/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.accessToken)

	resp, err := s.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	var u User
	err = json.NewDecoder(bytes.NewReader(buf)).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetAccounts retrieves all those bank accounts the user has chosen to share with your application
func (s *Session) GetAccounts() ([]*Account, error) {
	req, err := http.NewRequest("GET", s.Endpoint+"/rest/accounts", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.accessToken)

	resp, err := s.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	var answer struct {
		Accounts []*Account `json:'accounts'`
	}
	return answer.Accounts, json.NewDecoder(bytes.NewReader(buf)).Decode(&answer)
}
