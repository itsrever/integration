package test

import (
	"encoding/json"
	"os"
)

// Config is the configuration of the suite of tests
type Config struct {
	// BaseURL is the host or host + fixed part of the path (for every method)
	BaseURL string `json:"base_url"`
	// Auth is the authentication information (api key only for now)
	Auth *AuthenticationInfo `json:"auth"`
	// Tests is the configuration of tests to run
	Tests []Test `json:"tests"`
}

type AuthenticationInfo struct {
	HeaderName string `json:"header"`
	ApiKey     string `json:"api_key"`
}

type Test struct {
	Method     string     `json:"method"`
	UrlPattern string     `json:"url_pattern"`
	Scenarios  []Scenario `json:"scenarios"`
}

type Scenario struct {
	Name                   string `json:"name"`
	CustomerPrintedOrderId string `json:"customer_printed_order_id"`
	Optional               bool   `json:"optional"`
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

func (s *Scenario) Vars() map[string]string {
	return map[string]string{
		"customer_printed_order_id": s.CustomerPrintedOrderId,
		"optional":                  boolToString(s.Optional),
	}
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
