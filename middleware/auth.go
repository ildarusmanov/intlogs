package middleware

import (
	"net/http"
)

type Auth struct {
	AuthToken string
}

func CreateNewAuth(authToken string) Auth {
	return Auth{authToken}
}

func (a Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isValidToken := a.AuthToken == r.URL.Query().Get("token")

	if (!isValidToken) {
		panic("Invalid auth token")
	}
}
