package services

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewFilters(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	assert := assert.New(t)
	assert.NotNil(CreateNewFilters(ctx))
}

func TestParseQuery(t *testing.T) {
	// todo: add tests
}
