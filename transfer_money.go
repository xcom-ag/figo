package figo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Session) CreatePayment(p *Payment) (*Payment, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", s.Endpoint+"/rest/accounts/"+p.AccountID+"/payments", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.accessToken)

	resp, err := s.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	var n Payment
	return &n, json.NewDecoder(bytes.NewReader(buf)).Decode(&n)
}

func (s *Session) SubmitPayment(accountID, paymentID, TANSchemeID string) (string, error) {
	r := struct {
		RedirectURI string `json:"redirect_uri"`
		State       string `json:"state"`
		TANSchemeID string `json:"tan_scheme_id"`
	}{
		TANSchemeID: TANSchemeID,
	}
	data, err := json.Marshal(&r)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", s.Endpoint+"/rest/accounts/"+accountID+"/payments/"+paymentID+"/submit", bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.accessToken)

	resp, err := s.hc.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	log.Printf("post payment response: %s", string(buf))
	var a struct {
		TaskToken string `json:"task_token"`
	}
	err = json.NewDecoder(bytes.NewReader(buf)).Decode(&a)
	return a.TaskToken, err
}
