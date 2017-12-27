package user

import (
	"net/http/httptest"
	"net/http"
	"bytes"
	"testing"
)

var authToken = "test-token"

func createHttpRequestExample(method, addr, token string, body string) *http.Request {
	req := httptest.NewRequest(
		method,
		addr,
		bytes.NewBufferString(body),
	)

	SetAuthTokenHeader(token, req)

	return req
}

func TestCreateNewAuth(t *testing.T) {
	auth := CreateNewAuth(authToken)

	if auth == nil {
		t.Error("Auth does not created")
	}

	if auth.GetToken() != authToken {
		t.Error("Invalid token value")
	}
}

func TestValidateMethod(t *testing.T) {
	auth := CreateNewAuth(authToken)

	validTokenReq := createHttpRequestExample(
		"POST",
		"http://test.com/test_url.html",
		authToken,
		"{}",
	)

	if !auth.ValidateRequest(validTokenReq) {
		t.Error("Correct token does not accepted")
	}

	invalidTokenReq := createHttpRequestExample(
		"POST",
		"http://test.com/test_url.html",
		"incorrect token",
		"{}",
	)

	if auth.ValidateRequest(invalidTokenReq) {
		t.Error("Incorrect token accepted")
	}
}
