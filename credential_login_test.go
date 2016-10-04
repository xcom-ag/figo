package figo_test

import (
	"testing"

	"github.com/xcom-ag/figo"
)

func TestCredentialLogin(t *testing.T) {
	api := figo.NewAPI(DemoClientID, DemoClientSecret)
	apiToken, err := api.CredentialLogin(DemoUserLogin, DemoUserPassword)
	if err != nil {
		t.Logf("failed credentail login: %v", err)
		// fail always, because demo account got an endless valid access token per definition
		// closed without notice in API doc: https://github.com/figo-connect/node-figo/issues/17
		return
	}
	t.Logf("new API token: %s", apiToken)
}
