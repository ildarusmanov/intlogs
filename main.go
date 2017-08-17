package main

import (
	"intlogs/configs"
	"intlogs/middlewares"
	"fmt"

	"github.com/WajoxSoftware/middleware"
)

func main() {
	fmt.Println("Starting application...")

	fmt.Println("Load config")
	config := configs.LoadConfigFile()

	fmt.Println("Open MongoDB session")
	mgoSession := CreateMgoSession(config.MgoUrl)
	defer mgoSession.Close()

	fmt.Println("Create router")
	router := CreateNewRouter(mgoSession, config)

	fmt.Println("Define middleware")
	mware := middleware.CreateNewMiddleware()
	mware.AddHandler(middlewares.CreateNewAuth(config.AuthToken))
	mware.AddHandler(middlewares.CreateNewJsonOkResponse())
	mware.AddHandler(router)

	fmt.Println("Start web-server")
	StartServer(mware, config)
}