package main

import (
	"intlogs/configs"
	"intlogs/controllers"
	"log"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/mux"
	"net/http"
)

type RouterHandler struct {
	router *mux.Router
}

func CreateNewRouterHandler(dbSession *mgo.Session, config *configs.Config) *RouterHandler {
	return &RouterHandler{createNewRouter(dbSession, config)}
}

func (h *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	h.router.ServeHTTP(w,r)

	return true
}

func createNewRouter(dbSession *mgo.Session, config *configs.Config) *mux.Router {
	router := mux.NewRouter()

	log.Printf("Create controller")
	controller := controllers.CreateNewActionLogController(dbSession, config)

	log.Printf("Define routes")
	router.HandleFunc("/create", controller.CreateHandler).Methods("POST")
	router.HandleFunc("/get", controller.IndexHandler).Methods("GET")

	return router
}

