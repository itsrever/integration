package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// Config is the configuration of the suite of tests
type Config struct {
	// BaseURL is the host or host + fixed part of the path (for every method)
	BaseURL string `json:"base_url"`
	// ApiKeyAuth is the authentication information for Api Key authentication
	ApiKeyAuth *ApiKeyAuthInfo `json:"auth"`
	// OAuth2Info is the authentication information for OAuth2 authentication
	OAuth2Info *OAuth2Info `json:"oauth2"`
	// Tests is the configuration of tests to run
	Tests []Test `json:"tests"`
	// Debug the requests/responses
	Debug bool `json:"debug"`
}

// ApiKeyAuthInfo is the configuration for API key authentication
type ApiKeyAuthInfo struct {
	HeaderName string `json:"header"`
	ApiKey     string `json:"api_key"`
}

// OAuth2Info is the configuration for OAuth2 authentication
type OAuth2Info struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	// Url or Path in where the API is serving the OAuth2 token generation endpoint
	TokenUrl  string   `json:"token_url"`
	Scopes    []string `json:"scopes"`
	GrantType string   `json:"grant_type"`
}

type Test struct {
	Method     string     `json:"method"`
	UrlPattern string     `json:"url_pattern"`
	Scenarios  []Scenario `json:"scenarios"`
}

type Scenario struct {
	Name                   string `json:"name"`
	CustomerPrintedOrderId string `json:"customer_printed_order_id"`
	OrderID                string `json:"order_id"`
}

func configFromEnv() (*Config, error) {
	pathConfig := "config.json"
	overridePath := os.Getenv("TEST_CONFIG")
	if overridePath != "" {
		pathConfig = overridePath
	}
	return NewConfig(pathConfig)
}

func NewConfig(path string) (*Config, error) {
	configBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = json.Unmarshal(configBytes, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Test finds a test by method name
func (c *Config) Test(method string) *Test {
	for _, test := range c.Tests {
		if test.Method == method {
			return &test
		}
	}
	return nil
}

// Scenario finds a scenario by name
func (t *Test) Scenario(name string) *Scenario {
	for _, scenario := range t.Scenarios {
		if scenario.Name == name {
			return &scenario
		}
	}
	return nil
}

// SkipIfNotPresent skips the test if the scenario is not present in the config
func (t *Test) SkipTestIfScenarioNotPresent(tt *testing.T, scenarioName string) {
	found := false
	for _, scenario := range t.Scenarios {
		found = scenario.Name == scenarioName
		if found {
			break
		}
	}
	if !found {
		tt.Skipf("Skipping: Scenario %v not present in config", scenarioName)
	}
}

// FailTestIfScenarioNotPresent requires the test config to be present
func (t *Test) FailTestIfScenarioNotPresent(tt *testing.T, scenarioName string) {
	found := false
	for _, scenario := range t.Scenarios {
		found = scenario.Name == scenarioName
		if found {
			break
		}
	}
	if !found {
		//nolint:gosimple
		tt.Fatal(fmt.Sprintf("Error: Scenario %v not present in config", scenarioName))
	}
}

func (s *Scenario) Vars() map[string]string {
	return map[string]string{
		"customer_printed_order_id": s.CustomerPrintedOrderId,
		"order_id":                  s.OrderID,
	}
}
