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

	if cost := f.getIntValFromContext("cost"); cost != 0 {
		filters["action_cost"] = cost
	}

	return filters
}

func (f *Filters) getCreatedFilters() (interface{}, error) {
	createdFrom := f.getIntValFromContext("createdFrom")
	createdTo := f.getIntValFromContext("createdTo")

	if createdFrom == 0 && createdTo == 0 {
		return nil, errors.New("No filters")
	}
	createdFitlers := bson.M{}

	if createdFrom != 0 {
		createdFitlers["$gte"] = createdFrom
	}

	if createdTo != 0 {
		createdFitlers["$lte"] = createdTo
	}

	return createdFitlers, nil
}

func (f *Filters) getCostFilters() (interface{}, error) {
	costFrom := f.getIntValFromContext("costFrom")
	costTo := f.getIntValFromContext("costTo")
	cost := f.getIntValFromContext("cost")

	if cost != 0 {
		return cost, nil
	}

	if costFrom == 0 && costTo == 0 {
		return nil, errors.New("No filters")
	}

	costFitlers := bson.M{}

	if costFrom != 0 {
		costFitlers["$gte"] = costFrom
	}

	if costTo != 0 {
		costFitlers["$lte"] = costTo
	}

	return costFitlers, nil
}

func (f *Filters) getStrFromContext(key string) string {
	return f.context.Request.URL.Query().Get(key)
}

func (f *Filters) getIntValFromContext(key string) int {
	str := f.getStrFromContext(key)

	if str == "" {
		return 0
	}

	if i, err := strconv.Atoi(str); err == nil {
		return i
	}

	return 0
}
