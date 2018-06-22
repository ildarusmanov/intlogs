package user

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert := assert.New(t)
	assert.NotNil(auth)
	assert.Equal(auth.GetToken(), authToken)
}

func TestValidateMethod(t *testing.T) {
	auth := CreateNewAuth(authToken)

	validTokenReq := createHttpRequestExample(
		"POST",
		"http://test.com/test_url.html",
		authToken,
		"{}",
	)

	invalidTokenReq := createHttpRequestExample(
		"POST",
		"http://test.com/test_url.html",
		"incorrect token",
		"{}",
	)

	assert := assert.New(t)
	assert.True(auth.ValidateRequest(validTokenReq))
	assert.False(auth.ValidateRequest(invalidTokenReq))
}
