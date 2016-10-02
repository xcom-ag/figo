package figo_test

import (
	"testing"

	"github.com/xcom-ag/figo"
)

func TestGetAccounts(t *testing.T) {
	s := figo.NewSession(DemoAccessToken)

	accounts, err := s.GetAccounts()
	if err != nil {
		t.Errorf("get user failed: %v", err)
		return
	}

	for _, acct := range accounts {
		t.Logf("got account: %s (%s): %s %s", acct.AccountID, acct.Name, acct.IBAN, acct.BIC)
	}

}
