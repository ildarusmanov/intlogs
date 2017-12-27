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

func CreateNewRouter(dbSession *mgo.Session, config *configs.Config) *gin.Engine {
	r := gin.Default()

	log.Printf("Create controller")
	controller := controllers.CreateNewActionLogController(dbSession, config)

	log.Printf("Define routes")


	v1 := r.Group("/v1")
	{
		v1.Use(func(c *gin.Context) {
			if !user.CreateNewAuth(config.AuthToken).ValidateRequest(c.Request) {
				c.JSON(http.StatusForbidden, gin.H{"error": "invalid token"})
			} else {
				c.Next()
			}
		})

		v1.GET("/get", controller.IndexHandler)
		v1.POST("/create", controller.CreateHandler)
	}

	return r
}
