package controllers

import (
	"intlogs/configs"
	"intlogs/models"
	"intlogs/stores"
	"net/http"
	"strconv"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"gopkg.in/validator.v2"
)

const (
	PAGE_SIZE int = 0
	COLLECTION_NAME string = "action_logs"
)

type ActionLogController struct {
	store *stores.ActionLogStore
}

func CreateNewActionLogController(dbSession *mgo.Session,config *configs.Config) *ActionLogController {
	collection := dbSession.DB(config.MgoDb).C(COLLECTION_NAME)
	store := stores.CreateNewActionLogStore(collection)

	return &ActionLogController{store}
}

func (c *ActionLogController) CreateHandler(w http.ResponseWriter, r *http.Request) {
    log := models.CreateNewActionLog()

    if err := json.NewDecoder(r.Body).Decode(log); err != nil {
   		panic(err)
    }

	if err := validator.Validate(log); err != nil {
		panic(err)
	}

 	c.store.Save(log)

    if err := json.NewEncoder(w).Encode(log); err != nil {
        panic(err)
    }
}

func (c *ActionLogController) IndexHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		page = 0
	}

	logs := models.CreateNewActionLogCollection()
	limit := PAGE_SIZE * page
	
	c.store.All(logs, PAGE_SIZE, limit)

    if err := json.NewEncoder(w).Encode(logs); err != nil {
        panic(err)
    }
}
