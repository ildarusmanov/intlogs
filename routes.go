package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/intlogs/configs"
	"github.com/ildarusmanov/intlogs/controllers"
	"github.com/ildarusmanov/intlogs/user"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func createNewRouter(dbSession *mgo.Session, config *configs.Config) *gin.Engine {
	r := gin.Default()

	log.Printf("Create controller")
	controller := controllers.CreateNewActionLogController(dbSession, config)

	log.Printf("Define routes")

	r.Use(func(c *gin.Context) {
		if !user.CreateNewAuth(config.AuthToken).ValidateRequest(c.Request) {
			c.JSON(http.StatusForbidden, gin.H{"error": "invalid token"})
		}

		c.Next()
	})

	r.GET("/get", controller.IndexHandler)
	r.POST("/create", controller.CreateHandler)

	return router
}
