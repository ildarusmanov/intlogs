package controllers

import (
	"intlogs/configs"
	"intlogs/models"

	"net/http"

	"gopkg.in/mgo.v2"
	"github.com/gorilla/schema"
	"encoding/json"
)

type ActionLogController struct {
	mgoSession *mgo.Session
	config *configs.Config
}

func CreateNewActionLogController(mgoSession *mgo.Session,config *configs.Config) *ActionLogController {
	return &ActionLogController{mgoSession, config}
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

 	/**
 	 * @todo add save logic
 	 */
 
    if err = json.NewEncoder(w).Encode(log); err != nil {
        panic(err)
    }
}

func (c *ActionLogController) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	logs := models.CreateNewActionLogCollection()

    if err := json.NewEncoder(w).Encode(logs); err != nil {
        panic(err)
    }
}
