package controllers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/ildarusmanov/intlogs/db"
	"github.com/ildarusmanov/intlogs/models"
	"github.com/ildarusmanov/intlogs/tests"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestIndexHandler(t *testing.T) {
	config := tests.CreateConfig()
	dbSession := db.CreateSession(config.MgoUrl)
	defer dbSession.Close()

	controller := CreateNewActionLogController(dbSession, config)

	router := gin.New()
	router.GET("/get", controller.IndexHandler)

	req := httptest.NewRequest("GET", "http://127.0.0.1:8000/get?page=0", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	body := []byte(resp.Body.String())
	logs := models.MakeNewActionLogCollection()

	err := json.Unmarshal(body, &logs)

	assert.Nil(t, err)
}

func TestCreateHandler(t *testing.T) {
	config := tests.CreateConfig()
	dbSession := db.CreateSession(config.MgoUrl)
	defer dbSession.Close()

	controller := CreateNewActionLogController(dbSession, config)

	router := gin.New()
	router.POST("/create", controller.CreateHandler)

	bodyJson := "{\"ActionName\": \"authorized\", \"ActionTarget\": \"user\", \"ActionCost\": 1000, \"UserId\": \"some-user-id\", \"GuestUserId\": \"some-guest-id\", \"Url\": \"http://test.com\", \"CreatedAt\": 1712311}"
	inBody := bytes.NewBufferString(bodyJson)

	req := httptest.NewRequest("POST", "http://127.0.0.1:8000/create", inBody)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	body := []byte(resp.Body.String())
	log := models.CreateNewActionLog()

	err := json.Unmarshal(body, log)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(log.ActionName, "authorized")
}
