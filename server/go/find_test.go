package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindResponses(t *testing.T) {
	assert.Nil(t, FindResponseFor("non-existing-order"))
	assert.NotNil(t, FindResponseFor("simple_order_1"))
	assert.NotNil(t, FindResponseFor("simple_order_2"))
}
