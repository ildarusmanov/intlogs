package middleware

import (
	"net/http"
)

type MiddlewareHandler interface {
	Exec(w http.ResponseWriter, r *http.Request) bool
}

type MiddlewareHandlers []http.MiddlewareHandler

func CreateNewMiddlewareHandlers() *MiddlewareHandlers {
	return &MiddlewareHandlers{}
}

func (mHandlers *MiddlewareHandlers) GetHandler(finalH http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, h := range mHandlers {
			if (!h.Exec(w, r)) {
				return
			}
		}

		finalH.serveHTTP(w, r)
	})
}
