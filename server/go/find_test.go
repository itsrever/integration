package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindResponses(t *testing.T) {
	assert.Nil(t, FindOrderFor("non-existing-order"))
	assert.NotNil(t, FindOrderFor("simple_order_1"))
	assert.NotNil(t, FindOrderFor("simple_order_2"))
}
