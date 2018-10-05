package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ildarusmanov/intlogs/models"
)

var _ = Describe("Actionlogcollection", func() {
	var (
		collection *ActionLogCollection
	)

	BeforeEach(func() {
		collection = CreateNewActionLogCollection()
	})

    Describe("Create collection", func() {
        Context("Empty collection", func() {
            It("Should not be empty", func () {
				Expect(collection).NotTo(BeNil())
			})
        })
    })
})
