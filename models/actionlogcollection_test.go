package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewActionLogCollection(t *testing.T) {
	collection := CreateNewActionLogCollection()

	assert.NotNil(t, collection)
}
