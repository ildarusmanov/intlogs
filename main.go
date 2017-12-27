package main

import (
	"github.com/ildarusmanov/intlogs/configs"
	"github.com/ildarusmanov/intlogs/db"
	"github.com/ildarusmanov/intlogs/middlewares"

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
	routerHandler := CreateNewRouterHandler(dbSession, config)
	fmt.Println("Start web-server")
	StartServer(routerHandler, config)
}
