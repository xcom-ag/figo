package figo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
	buf, _ := ioutil.ReadAll(resp.Body)
	var answer struct {
		Accounts []*Account `json:"accounts"`
	}
	return answer.Accounts, json.NewDecoder(bytes.NewReader(buf)).Decode(&answer)
}
