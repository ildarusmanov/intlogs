package services

import (
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

	if createdFrom := f.getIntValFromContext("createdFrom"); createdFrom != 0 {
		filters["created_at"] = bson.M{"$gte": createdFrom}
	}

	if createdTo := f.getIntValFromContext("createdTo"); createdTo != 0 {
		filters["created_at"] = bson.M{"$lte": createdTo}
	}

	if costFrom := f.getIntValFromContext("costFrom"); costFrom != 0 {
		filters["action_cost"] = bson.M{"$gte": costFrom}
	}

	if costTo := f.getIntValFromContext("costTo"); costTo != 0 {
		filters["action_cost"] = bson.M{"$gte": costTo}
	}

	if cost := f.getIntValFromContext("cost"); cost != 0 {
		filters["action_cost"] = cost
	}

	return filters
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
