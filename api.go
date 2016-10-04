package figo

import "net/http"

type API struct {

	// Endpoint of figo connect API - usually set with BaseAPIEndpoint (https://api.figo.me)
	Endpoint string

	clientID     string
	clientSecret string
	hc           *http.Client
}

func NewAPI(clientID, clientSecret string) *API {
	return &API{
		Endpoint:     BaseAPIEndpoint,
		clientID:     clientID,
		clientSecret: clientSecret,
		hc:           &http.Client{},
	}
}
