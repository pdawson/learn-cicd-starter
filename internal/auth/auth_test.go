package auth_test

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{"Authorization": {"ApiKey Test"}}

	_, e := auth.GetAPIKey(headers)

	if e != nil {
		t.Fatalf("%v", e)
	}
}
