package controllers

import (
	"intlogs/configs"
	"intlogs/models"
	"intlogs/stores"

	"net/http"
	"strconv"

	"gopkg.in/mgo.v2"
	"github.com/gorilla/schema"
	"encoding/json"
)

const (
	PAGE_SIZE int = 0
)

type ActionLogController struct {
	store *stores.ActionLogStore
}

func CreateNewActionLogController(mgoSession *mgo.Session,config *configs.Config) *ActionLogController {
	collection := mgoSession.DB(config.MgoDb).C("action_logs")
	store := stores.CreateNewActionLogStore(collection)

	return &ActionLogController{store}
}

func (c *ActionLogController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

    err := r.ParseForm()

    if err != nil {
   		panic(err)
    }

    log := models.CreateNewActionLog()

    decoder := schema.NewDecoder()
    // r.PostForm is a map of our POST form values
    err = decoder.Decode(log, r.PostForm)

    if err != nil {
   		panic(err)
    }

 	c.store.Save(log)

    if err = json.NewEncoder(w).Encode(log); err != nil {
        panic(err)
    }
}

func (c *ActionLogController) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if (err != nil) {
		page = 0
	}

	logs := models.CreateNewActionLogCollection()
	limit := PAGE_SIZE * page
	
	c.store.All(logs, PAGE_SIZE, limit)

    if err := json.NewEncoder(w).Encode(logs); err != nil {
        panic(err)
    }
}
