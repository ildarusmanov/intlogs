package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/intlogs/configs"
	"github.com/ildarusmanov/intlogs/controllers"
	"github.com/ildarusmanov/intlogs/user"
	"gopkg.in/mgo.v2"
)

func CreateNewRouter(dbSession *mgo.Session, config *configs.Config) *gin.Engine {
	r := gin.Default()

	log.Printf("Create controller")

	actionlogs := controllers.CreateNewActionLogController(dbSession, config)
	status := controllers.CreateNewStatusController()

	log.Printf("Define routes")

	v1 := r.Group("/api/v1")
	{
		v1.Use(func(c *gin.Context) {
			if !user.CreateNewAuth(config.AuthToken).ValidateRequest(c.Request) {
				c.JSON(http.StatusForbidden, gin.H{"error": "invalid token"})
			} else {
				c.Next()
			}
		})

		v1.GET("/get", actionlogs.IndexHandler)
		v1.POST("/create", actionlogs.CreateHandler)
		v1.GET("/status", status.CheckHandler)
	}

	return r
}
