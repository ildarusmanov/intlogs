package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
	"testing"
	"time"
)

func TestValidators(t *testing.T) {
	invalidLog := CreateNewActionLog()

	if err := validator.Validate(invalidLog); err == nil {
		t.Error("Empty ActionLog model validation: Error expected, but", err, "given")
	}

	validLog := CreateNewActionLog()
	validLog.ActionName = "authorized"
	validLog.ActionTarget = "user"
	validLog.ActionCost = 1000
	validLog.UserId = "some-user-id"
	validLog.GuestUserId = "some-guest-id"
	validLog.Url = "http://test.com"
	validLog.CreatedAt = time.Now().Unix()
	validLog.Params = bson.M{"key1": "value1"}

	if err := validator.Validate(validLog); err != nil {
		t.Error("Valid ActionLog model validation: Nil expected, but", err, "given")
	}
}