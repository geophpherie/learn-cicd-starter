package auth

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestNoHeader(t *testing.T) {
	headers := http.Header(map[string][]string{})
	_, err := GetAPIKey(headers)
	if errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fail()
	}
}

func TestValidHeader(t *testing.T) {
	expectedApiKey := "myapikey"
	headers := http.Header(map[string][]string{})
	headers.Add("Authorization", fmt.Sprintf("ApiKey %v", expectedApiKey))

	parsedApiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("%v", err)
	}
	if parsedApiKey != expectedApiKey {
		t.Errorf("%v is not %v", parsedApiKey, expectedApiKey)
	}
}
