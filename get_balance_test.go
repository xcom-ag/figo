package figo_test

import (
	"testing"

	"github.com/xcom-ag/figo"
)

func TestGetBalance(t *testing.T) {
	s := figo.NewSession(DemoAccessToken)

	bal, err := s.GetBalance("A1.1")
	if err != nil {
		t.Errorf("get balance failed: %v", err)
		return
	}
	t.Logf("got balance: %f (%s)", bal.Balance, bal.BalanceDate)

}
