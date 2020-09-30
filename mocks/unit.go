package mocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	total := 0
	// total := Add(1, 2)
	assert.Equal(t, total, 2)

}

//notnil
func getDBconnection(t *testing.T) {
	client := 0
	assert.NotNil(t, client, "Not nil")
}

// https://godoc.org/github.com/stretchr/testify/assert
