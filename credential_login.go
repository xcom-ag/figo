package figo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (api *API) CredentialLogin(username, password string) (string, error) {
	r := struct {
		GrantType string `json:"grant_type"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}{
		GrantType: "password",
		Username:  username,
		Password:  password,
	}
	data, err := json.Marshal(&r)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", api.Endpoint+"/auth/token", bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(api.clientID, api.clientSecret)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := api.hc.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	var a struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
	}
	err = json.NewDecoder(bytes.NewReader(buf)).Decode(&a)
	return a.AccessToken, err
}
