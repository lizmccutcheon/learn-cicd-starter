package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey_HeaderMissing(t *testing.T) {
	header := http.Header{}

	_, err := GetAPIKey(header)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("got error: %v, want: %v", err, ErrNoAuthHeaderIncluded)
	}
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "Bearer this-is-wrong")

	_, err := GetAPIKey(header)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	expectedErrMsg := "malformed authorization header"
	if err.Error() != expectedErrMsg {
		t.Errorf("got error message: %q, want: %v", err.Error(), expectedErrMsg)
	}
}

func TestGetAPIKey_Success(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey fake-test-key")

	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "fake-test-key"
	if got != want {
		t.Errorf("got API key: %q, want: %q", got, want)
	}

}
