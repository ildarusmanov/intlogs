package models_test

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ildarusmanov/intlogs/models"
)

var _ = Describe("Actionlog", func() {
	var (
		actionName     string
		actionTarget   string
		actionTargetId string
		actionCost     int64
		userId         string
		guestUserId    string
		url            string
		createdAt      int64
		params         bson.M

		emptyModel *ActionLog
		validModel *ActionLog
	)

	BeforeEach(func() {
		actionName = "authorized"
		actionTarget = "user"
		actionTargetId = "1"
		actionCost = 1000
		userId = "some-user-id"
		guestUserId = "some-guest-id"
		url = "http://test.com"
		createdAt = time.Now().Unix()
		params = bson.M{"key1": "value1"}

		emptyModel = CreateNewActionLog()
		validModel = CreateNewActionLog()

		validModel.ActionName = actionName
		validModel.ActionTarget = actionTarget
		validModel.ActionTargetId = actionTargetId
		validModel.ActionCost = actionCost
		validModel.UserId = userId
		validModel.GuestUserId = guestUserId
		validModel.Url = url
		validModel.CreatedAt = createdAt
		validModel.Params = params
	})

	Describe("Empty ActionLog", func() {
		Context("Empty model", func() {
			It("Should not be nil", func() {
				Expect(emptyModel).NotTo(BeNil())
			})

			It("Should not be valid", func() {
				Expect(validator.Validate(emptyModel)).NotTo(BeNil())
			})
		})
	})

	Describe("Not Empty ActionLog", func() {
		Context("Filled model", func() {
			It("Should have all attributes", func() {
				Expect(validModel).NotTo(BeNil())
				Expect(validModel.ActionName).To(Equal(actionName))
				Expect(validModel.ActionTarget).To(Equal(actionTarget))
				Expect(validModel.ActionTargetId).To(Equal(actionTargetId))
				Expect(validModel.ActionCost).To(Equal(actionCost))
				Expect(validModel.UserId).To(Equal(userId))
				Expect(validModel.GuestUserId).To(Equal(guestUserId))
				Expect(validModel.Url).To(Equal(url))
				Expect(validModel.CreatedAt).To(Equal(createdAt))
				Expect(validModel.Params).To(Equal(params))
			})

			It("Should be valid", func() {
				Expect(validator.Validate(validModel)).To(BeNil())
			})
		})
	})
})
