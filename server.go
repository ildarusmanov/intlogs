package main

import (
	"intlogs/configs"

	"net/http"
	"log"
	"time"

	"github.com/gorilla/mux"
)

func StartServer(router *mux.Router, config *configs.Config) {
    srv := &http.Server{
        Handler:      router,
        Addr:         config.ServerHost,
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}