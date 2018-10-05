package services_test

import (
	"github.com/gin-gonic/gin"
	. "github.com/ildarusmanov/intlogs/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http/httptest"
)

var _ = Describe("Filters", func() {
	var (
		w       *httptest.ResponseRecorder
		ctx     *gin.Context
		filters *Filters
	)

	BeforeEach(func() {
		w = httptest.NewRecorder()
		ctx, _ = gin.CreateTestContext(w)
		filters = CreateNewFilters(ctx)
	})

	Describe("Create filters from context", func() {
		Context("With test context", func() {
			It("should be not nil", func() {
				Expect(filters).NotTo(BeNil())
			})
		})
	})

	Describe("Parse query", func() {
		Context("With new filters", func() {
			PIt("should parse empty filters", func() {})
		})
	})
})
