package services

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type Filters struct {
	context *gin.Context
}

func CreateNewFilters(ctx *gin.Context) *Filters {
	return &Filters{context: ctx}
}

func (f *Filters) ParseQuery() interface{} {
	filters := bson.M{}

	if userID := f.getStrFromContext("userId"); userID != "" {
		filters["user_id"] = userID
	}

	if guestUserID := f.getStrFromContext("guestUserId"); guestUserID != "" {
		filters["guest_user_id"] = guestUserID
	}

	if name := f.getStrFromContext("name"); name != "" {
		filters["action_name"] = name
	}

	if target := f.getStrFromContext("target"); target != "" {
		filters["action_target"] = target
	}

	if costFilters, err := f.getCostFilters(); err == nil {
		filters["action_cost"] = costFilters
	}

	if createdFilters, err := f.getCreatedFilters(); err == nil {
		filters["created_at"] = createdFilters
	}

	return filters
}

func (f *Filters) getCreatedFilters() (interface{}, error) {
	createdFrom, errFrom := f.getIntValFromContext("createdFrom")
	createdTo, errTo := f.getIntValFromContext("createdTo")

	if errFrom != nil && errTo != nil {
		return nil, errors.New("No filters")
	}

	createdFitlers := bson.M{}

	if errFrom == nil {
		createdFitlers["$gte"] = createdFrom
	}

	if errTo == nil {
		createdFitlers["$lte"] = createdTo
	}

	return createdFitlers, nil
}

func (f *Filters) getCostFilters() (interface{}, error) {
	costFrom, errFrom := f.getIntValFromContext("costFrom")
	costTo, errTo := f.getIntValFromContext("costTo")
	cost, err := f.getIntValFromContext("cost")

	if err == nil {
		return cost, nil
	}

	if errFrom != nil && errTo != nil {
		return nil, errors.New("No filters")
	}

	costFilters := bson.M{}

	if errFrom == nil {
		costFilters["$gte"] = costFrom
	}

	if errTo == nil {
		costFilters["$lte"] = costTo
	}

	return costFilters, nil
}

func (f *Filters) getStrFromContext(key string) string {
	return f.context.Request.URL.Query().Get(key)
}

func (f *Filters) getIntValFromContext(key string) (int, error) {
	str := f.getStrFromContext(key)

	if str == "" {
		return 0, errors.New("Empty")
	}

	i, err := strconv.Atoi(str)

	if err != nil {
		return 0, err
	}

	return i, nil
}
