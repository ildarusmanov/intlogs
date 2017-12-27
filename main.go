package main

import (
	"intlogs/configs"
	"intlogs/db"
	"fmt"
	"path/filepath"
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
	routerHandler := CreateNewRouter(dbSession, config)
	fmt.Println("Start web-server")
	StartServer(routerHandler, config)
}
