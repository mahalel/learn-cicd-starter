package auth

import (
	"net/http"
	"testing"
)

func TestEmptyAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 12345")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if apiKey != "12345" {
		t.Errorf("GetAPIKey() = %q, want %q", apiKey, "12345")
	}
}
