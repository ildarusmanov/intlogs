package controllers

import (
	"intlogs/configs"
	"intlogs/models"
	"intlogs/stores"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/validator.v2"
	"net/http"
	"strconv"
)

const (
	PAGE_SIZE int = 20
)

type ActionLogController struct {
	store *stores.ActionLogStore
}

func CreateNewActionLogController(dbSession *mgo.Session, config *configs.Config) *ActionLogController {
	collection := dbSession.DB(config.MgoDb).C(config.MgoCollection)
	store := stores.CreateNewActionLogStore(collection)

	return &ActionLogController{store}
}

func (c *ActionLogController) CreateHandler(context *gin.Context) {
	log := models.CreateNewActionLog()

	if err := json.NewDecoder(context.Request.Body).Decode(log); err != nil {
		panic(err)
	}

	if err := validator.Validate(log); err != nil {
		panic(err)
	}

	if _, err := c.store.Save(log); err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, log)
}

func (c *ActionLogController) IndexHandler(context *gin.Context) {
	page, err := strconv.Atoi(context.Request.URL.Query().Get("page"))

	if err != nil {
		page = 0
	}

	logs := models.CreateNewActionLogCollection()
	offset := PAGE_SIZE * page

	if _, err := c.store.All(logs, PAGE_SIZE, offset); err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, logs)
}
