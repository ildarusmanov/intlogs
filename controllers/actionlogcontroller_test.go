package controllers

import (
	"github.com/ildarusmanov/intlogs/db"
	"github.com/ildarusmanov/intlogs/models"
	"github.com/ildarusmanov/intlogs/tests"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
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
	r := httptest.NewRequest("GET", "http://127.0.0.1:8000/get?page=0", inBody)

	controller.IndexHandler(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	logs := models.MakeNewActionLogCollection()

	if err := json.Unmarshal(body, &logs); err != nil {
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
	r := httptest.NewRequest("POST", "http://127.0.0.1:8000/create", inBody)

	controller.CreateHandler(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	log := models.CreateNewActionLog()

	if err := json.Unmarshal(body, log); err != nil {
		t.Error("Invalid json response")
	}

	if log.ActionName != "authorized" {
		t.Error("Incorrect data")
	}
}
