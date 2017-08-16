package middleware

import (
	"intlogs/configs"

	"net/http"
)

type Auth struct {
	AuthToken string
}

func CreateNewAuth(authToken string) Auth {
	return Auth{authToken}
}

func (a *Auth) Exec(w http.ResponseWriter, r *http.Request) bool {
	isValidToken := a.AuthToken == r.URL.Query().Get("token")

	return IsValidToken
}
