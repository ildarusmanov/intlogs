package controllers

import (
	"github.com/ildarusmanov/intlogs/configs"
	"github.com/ildarusmanov/intlogs/models"
	"github.com/ildarusmanov/intlogs/stores"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
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
	page := c.getPageFromContext(context)
	filters := c.getFiltersFromContext(context)

	logs := models.CreateNewActionLogCollection()
	offset := PAGE_SIZE * page

	if _, err := c.store.All(logs, filters, PAGE_SIZE, offset); err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, logs)
}

func (c *ActionLogController) getPageFromContext(context *gin.Context) int {
	pageStr := context.Request.URL.Query().Get("page")

	if pageStr == "" {
		return 0
	}

	if page, err := strconv.Atoi(pageStr); err == nil {
		return page
	}

	return 0
}

func (c *ActionLogController) getUserIdFromContext(context *gin.Context) string {
	return context.Request.URL.Query().Get("userId")
}

func (c *ActionLogController) getFiltersFromContext(context *gin.Context) interface{} {
	userId := c.getUserIdFromContext(context)

	if userId == "" {
		return nil
	}

	return bson.M{"user_id": userId}
}
