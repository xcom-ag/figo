package figo_test

import (
	"testing"

	"github.com/xcom-ag/figo"
)

func TestGetAccounts(t *testing.T) {
	s := figo.NewSession(DemoAccessToken)

	accounts, err := s.GetAccounts()
	if err != nil {
		t.Errorf("get accounts failed: %v", err)
		return
	}

	for _, acct := range accounts {
		t.Logf("got account: %s (%s): %s %s", acct.AccountID, acct.Name, acct.IBAN, acct.BIC)
		if acct.Balance != nil {
			bal := acct.Balance
			t.Logf("    with balance: %s %f (%s)", acct.Currency, bal.Balance, bal.BalanceDate)
		}
		for _, ts := range acct.SupportedTANSchemes {
			t.Logf("    with TANScheme: %s(%s): %s", ts.Name, ts.TANSchemeID, ts.MediumName)
		}
	}

}
