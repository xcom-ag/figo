package figo_test

import (
	"testing"

	"github.com/xcom-ag/figo"
)

func TestCreatePayment(t *testing.T) {
	s := figo.NewSession(DemoAccessToken)

	p := &figo.Payment{
		AccountID: "A1.1",
		Type:      "SEPA transfer",
		Name:      "Mr. Nobody",
		IBAN:      "DE40900900424711951501",
		Amount:    1.01,
		Purpose:   "test",
		Cents:     false,
	}

	p, err := s.CreatePayment(p)
	if err != nil {
		t.Errorf("create payment failed: %v", err)
		return
	}
	t.Logf("submitted payment with ID: %s", p.PaymentID)

	taskToken, err := s.SubmitPayment(p.AccountID, p.PaymentID, "M1.1")
	if err != nil {
		t.Errorf("submit payment failed: %v", err)
		return
	}
	t.Logf("submitted payment with task token: %s", taskToken)

}
