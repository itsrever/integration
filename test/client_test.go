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
	assert.Equal(t, applyVars("/orders/{orderId}", map[string]string{"orderId": "$&/"}), "/orders/%24%26%2F")
	assert.Equal(t, applyVars("/orders?id={orderId}&sort=ASC", map[string]string{"orderId": "$&/"}), "/orders?id=%24%26%2F&sort=ASC")
	assert.Equal(t, applyVars("/orders?id={orderId}&sort=ASC", map[string]string{"orderId": "123-345"}), "/orders?id=123-345&sort=ASC")
	assert.Equal(t, applyVars("/orders?id={orderId}&sort=ASC", map[string]string{"orderId": "123/345"}), "/orders?id=123%2F345&sort=ASC")
}

func TestComposeURL(t *testing.T) {
	assert.Equal(t, composeRequestURL("http://localhost:8080", "/orders/123"), "http://localhost:8080/orders/123")
	assert.Equal(t, composeRequestURL("http://localhost:8080/", "/orders/123"), "http://localhost:8080/orders/123")
	assert.Equal(t, composeRequestURL("http://localhost:8080/", "/orders/123"), "http://localhost:8080/orders/123")
	assert.Equal(t, composeRequestURL("http://localhost:8080/", "orders/123"), "http://localhost:8080/orders/123")
	assert.Equal(t, composeRequestURL("http://localhost:8080", "orders/123"), "http://localhost:8080/orders/123")
}

func clientFromConfig(cfg *Config) *Client {
	c := NewClient(cfg.BaseURL)
	if cfg.ApiKeyAuth != nil {
		c = c.WithAuth(cfg.ApiKeyAuth)
	}
	if cfg.OAuth2Info != nil {
		c = c.WithOAuth2(cfg.OAuth2Info)
	}
	if cfg.Debug {
		c = c.Debug()
	}
	return c
}
