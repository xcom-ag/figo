package figo_test

import (
	"testing"

	"github.com/xcom-ag/figo"
)

func TestGetUser(t *testing.T) {
	s := figo.NewSession(DemoAccessToken)
	user, err := s.GetUser()
	if err != nil {
		t.Errorf("get user failed: %v", err)
		return
	}
	t.Logf("got user: %s (%s)", user.UserID, user.Email)
}
