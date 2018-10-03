package services

import (
	"errors"
	"strconv"
	"strings"

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

	if userIDFilters, err := f.getUserIDFilters(); err == nil {
		filters["user_id"] = userIDFilters
	}

	if guestUserID := f.getStrFromContext("guestUserId"); guestUserID != "" {
		filters["guest_user_id"] = guestUserID
	}

	if nameFilters, err := f.getNameFilters(); err == nil {
		filters["action_name"] = nameFilters
	}

	if targetIDFilters, err := f.getTargetIDFilters(); err == nil {
		filters["action_target_id"] = targetIDFilters
	}

	if targetFilters, err := f.getTargetFilters(); err == nil {
		filters["action_target"] = targetFilters
	}

	if costFilters, err := f.getCostFilters(); err == nil {
		filters["action_cost"] = costFilters
	}

	if createdFilters, err := f.getCreatedFilters(); err == nil {
		filters["created_at"] = createdFilters
	}

	return filters
}

func (f *Filters) getNameFilters() (interface{}, error) {
	if name := f.getStrFromContext("name"); name != "" {
		return name, nil
	}

	if names := f.getStrFromContext("names"); names != "" {
		namesArr := strings.Split(names, ",")

		return bson.M{"$in": namesArr}, nil
	}

	return nil, errors.New("No filters")
}

func (f *Filters) getTargetFilters() (interface{}, error) {
	if target := f.getStrFromContext("target"); target != "" {
		return target, nil
	}

	if targets := f.getStrFromContext("targets"); targets != "" {
		targetsArr := strings.Split(targets, ",")

		return bson.M{"$in": targetsArr}, nil
	}

	return nil, errors.New("No filters")
}

func (f *Filters) getTargetIDFilters() (interface{}, error) {
	if targetID := f.getStrFromContext("targetId"); targetID != "" {
		return targetID, nil
	}

	if targetIDs := f.getStrFromContext("targetIds"); targetIDs != "" {
		targetIDsArr := strings.Split(targetIDs, ",")

		return bson.M{"$in": targetIDsArr}, nil
	}

	return nil, errors.New("No filters")
}

func (f *Filters) getUserIDFilters() (interface{}, error) {
	if userID := f.getStrFromContext("userId"); userID != "" {
		return userID, nil
	}

	if userIDs := f.getStrFromContext("userIds"); userIDs != "" {
		userIDsArr := strings.Split(userIDs, ",")

		return bson.M{"$in": userIDsArr}, nil
	}

	return nil, errors.New("No filters")
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
