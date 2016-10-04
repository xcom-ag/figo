package figo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetBalance
func (s *Session) GetBalance(accountID string) (*Balance, error) {
	req, err := http.NewRequest("GET", s.Endpoint+"/rest/accounts/"+accountID+"/balance", nil)
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
	var b Balance
	return &b, json.NewDecoder(bytes.NewReader(buf)).Decode(&b)
}
