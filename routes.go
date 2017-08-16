package main

import (
	"github.com/gorilla/mux"
)

func CreateNewRouter() *mux.Router {
	return mux.NewRouter()
}