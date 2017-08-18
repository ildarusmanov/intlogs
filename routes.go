package main

import (
	"intlogs/configs"
	"intlogs/controllers"

	"fmt"
	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
)

func CreateNewRouter(dbSession *mgo.Session, config *configs.Config) *mux.Router {
	router := mux.NewRouter()

	fmt.Println("Create controller")
	controller := controllers.CreateNewActionLogController(dbSession, config)

	fmt.Println("Define routes")
	router.HandleFunc("/create", controller.CreateHandler).Methods("POST")
	router.HandleFunc("/get", controller.IndexHandler).Methods("GET")

	return router
}
