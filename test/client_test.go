package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyVars(t *testing.T) {
	assert.Equal(t, applyVars("/orders/{orderId}", map[string]string{"orderId": "123"}), "/orders/123")
	assert.Equal(t, applyVars("/orders/{orderId}", map[string]string{"other": "123"}), "/orders/{orderId}")
	assert.Equal(t, applyVars("/orders/{orderId}{orderId}", map[string]string{"orderId": "123"}), "/orders/123123")
	assert.Equal(t, applyVars("/orders/{orderId}/{modelId}", map[string]string{"orderId": "123", "modelId": "abc"}), "/orders/123/abc")
}

func TestComposeURL(t *testing.T) {
	assert.Equal(t, composeRequestURL("http://localhost:8080", "/orders/123"), "http://localhost:8080/orders/123")
	assert.Equal(t, composeRequestURL("http://localhost:8080/", "/orders/123"), "http://localhost:8080/orders/123")
	assert.Equal(t, composeRequestURL("http://localhost:8080/", "/orders/123"), "http://localhost:8080/orders/123")
	assert.Equal(t, composeRequestURL("http://localhost:8080/", "orders/123"), "http://localhost:8080/orders/123")
	assert.Equal(t, composeRequestURL("http://localhost:8080", "orders/123"), "http://localhost:8080/orders/123")
}
