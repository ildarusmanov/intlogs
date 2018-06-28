package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
)

func TestValidators(t *testing.T) {
	invalidLog := CreateNewActionLog()
	errInvalid := validator.Validate(invalidLog)

	validLog := CreateNewActionLog()
	validLog.ActionName = "authorized"
	validLog.ActionTarget = "user"
	validLog.ActionTargetId = "1"
	validLog.ActionCost = 1000
	validLog.UserId = "some-user-id"
	validLog.GuestUserId = "some-guest-id"
	validLog.Url = "http://test.com"
	validLog.CreatedAt = time.Now().Unix()
	validLog.Params = bson.M{"key1": "value1"}
	errValid := validator.Validate(validLog)

	assert := assert.New(t)
	assert.NotNil(errInvalid)
	assert.Nil(errValid)
}
