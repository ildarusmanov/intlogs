package user

import (
	"net/http"
)

var tokenHeaderParam = "Auth-Token"

type Auth struct {
	authToken string
}

func CreateNewAuth(authToken string) *Auth {
	return &Auth{authToken}
}

func SetAuthTokenHeader(token string, req *http.Request) {
	req.Header.Set(tokenHeaderParam, token)
}

func (a Auth) ValidateRequest(req *http.Request) bool {
	return a.authToken == req.Header.Get(tokenHeaderParam)
}

func (a Auth) GetToken() string {
	return a.authToken
}
