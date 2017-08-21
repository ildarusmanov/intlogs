package main

import (
	"intlogs/configs"
	"intlogs/db"
	"intlogs/middlewares"

	"fmt"
	"path/filepath"

	"github.com/WajoxSoftware/middleware"
)

func main() {
	fmt.Println("Starting application...")

	fmt.Println("Load config")
	configFilePath, _ := filepath.Abs("./config.yml")
	config := configs.LoadConfigFile(configFilePath)

	fmt.Println("Open MongoDB session")
	dbSession := db.CreateSession(config.MgoUrl)
	defer dbSession.Close()

	fmt.Println("Create router")
	router := CreateNewRouter(dbSession, config)

	fmt.Println("Define middleware")
	mware := middleware.CreateNewMiddleware()
	mware.AddHandler(middlewares.CreateNewAuth(config.AuthToken))
	mware.AddHandler(middlewares.CreateNewJsonOkResponse())
	mware.AddHandler(router)

	fmt.Println("Start web-server")
	StartServer(mware, config)
}
