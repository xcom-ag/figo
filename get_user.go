package figo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
