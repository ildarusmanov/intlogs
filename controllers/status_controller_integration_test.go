package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateStatusController(t *testing.T) {
	controller := CreateNewStatusController()
	assert.NotNil(t, controller)
}

func TestCheckStatusController(t *testing.T) {
	controller := CreateNewStatusController()
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	controller.CheckHandler(ctx)

	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
}
