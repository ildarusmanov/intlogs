package main

import (
	"intlogs/configs"
	"intlogs/controllers"
	"fmt"
)

func main() {
	fmt.Println("Starting application...")

	fmt.Println("Load config")
	config := configs.LoadConfigFile()

	fmt.Println("Open MongoDB session")
	mgoSession := CreateMgoSession(config.MgoUrl)
	defer mgoSession.Close()

	fmt.Println("Create router")
	router := CreateNewRouter()

	fmt.Println("Create controller")
	controller := controllers.CreateNewActionLogController(mgoSession, config)

	fmt.Println("Define routes")
	router.HandleFunc("/create",controller.CreateHandler).Methods("POST")
	router.HandleFunc("/get", controller.IndexHandler).Methods("GET")

	fmt.Println("Start web-server")
	StartServer(router, config)
}