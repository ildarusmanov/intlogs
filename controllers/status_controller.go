package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusController provides health-check function
type StatusController struct{}

// CreateNewStatusController constructor for StatusController
func CreateNewStatusController() *StatusController {
	return &StatusController{}
}

// CheckHandler returns ok if the service is ok
func (c *StatusController) CheckHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
