package controllers

import (
	"intlogs/db"
	"intlogs/models"
	"intlogs/tests"

	"bytes"
	"testing"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestIndexHandler(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Error", r)
		}
	}()

	config := tests.CreateConfig()
	dbSession := db.CreateSession(config.MgoUrl)
	defer dbSession.Close()

	controller := CreateNewActionLogController(dbSession, config)

	inBody := bytes.NewBufferString("")
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/get", inBody)

	controller.IndexHandler(w, r)

	respBytes, _ := ioutil.ReadAll(r.Body)
	
	logs := models.CreateNewActionLogCollection()

    if err := json.NewDecoder(bytes.NewBuffer(respBytes)).Decode(logs); err != nil {
   		t.Error("Invalid json response")
    }
}

func TestCreateHandler(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Error", r)
		}
	}()

	config := tests.CreateConfig()
	dbSession := db.CreateSession(config.MgoUrl)
	defer dbSession.Close()

	controller := CreateNewActionLogController(dbSession, config)

	bodyJson := "{\"ActionName\": \"authorized\", \"ActionTarget\": \"user\", \"ActionCost\": 1000, \"UserId\": \"some-user-id\", \"GuestUserId\": \"some-guest-id\", \"Url\": \"http://test.com\", \"CreatedAt\": 1712311}"
	inBody := bytes.NewBufferString(bodyJson)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/create", inBody)

	controller.IndexHandler(w, r)


	respBytes, _ := ioutil.ReadAll(r.Body)
	
	log := models.CreateNewActionLog()

    if err := json.NewDecoder(bytes.NewBuffer(respBytes)).Decode(log); err != nil {
   		t.Error("Invalid json response")
    }

    if log.ActionName != "authorized" {
    	t.Error("Incorrect data")
    }
}
