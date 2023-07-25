package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Client is a very generic http client for making requests to url patterns
type Client struct {
	baseURL string
	client  *http.Client
	auth    *AuthenticationInfo
	debug   bool
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		// TODO: Add timeout
		client: &http.Client{},
		debug:  false,
	}
}

// WithAuth adds authentication information to the client
func (c *Client) WithAuth(auth *AuthenticationInfo) *Client {
	return &Client{
		baseURL: c.baseURL,
		client:  c.client,
		auth:    auth,
		debug:   c.debug,
	}
}

func (c *Client) WithNoAuth() *Client {
	return &Client{
		baseURL: c.baseURL,
		client:  c.client,
		auth:    nil,
		debug:   c.debug,
	}
}

// Debug activates debugging calls
func (c *Client) Debug() *Client {
	return &Client{
		baseURL: c.baseURL,
		client:  c.client,
		auth:    c.auth,
		debug:   true,
	}
}

// Do makes an HTTP request to the path pattern, with the vars context and body with a JSON body
// method is the verb of the request. Ex: GET, POST, PUT, DELETE
// pathPattern is a path pattern with variables in the form of {var_name}. Ex: /orders/{order_id}
// vars is a map of variables to be applied to the path pattern. Ex: map[string]string{"order_id": "123"}
func (c *Client) Do(method string, pathPattern string, vars map[string]string, body any) (resp *http.Response, err error) {
	var req *http.Request
	bodyBuf := &bytes.Buffer{}
	path := composeRequestURL(c.baseURL, applyVars(pathPattern, vars))

	if body != nil {
		err = json.NewEncoder(bodyBuf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	if bodyBuf.Len() > 0 {
		req, err = http.NewRequest(method, path, bodyBuf)
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, path, nil)
	}
	if err != nil {
		return nil, err
	}
	if c.auth != nil {
		req.Header.Set(c.auth.HeaderName, c.auth.ApiKey)
	}
	return c.do(req)
}

func composeRequestURL(baseURL string, path string) string {
	path = strings.TrimPrefix(path, "/")
	baseURL = strings.TrimSuffix(baseURL, "/")
	return fmt.Sprintf("%v/%v", baseURL, path)
}

func applyVars(pathPattern string, vars map[string]string) string {
	path := pathPattern
	for varName, varValue := range vars {
		path = strings.ReplaceAll(path, "{"+varName+"}", url.QueryEscape(varValue))
	}
	return path
}

// do the request internally
func (c *Client) do(request *http.Request) (*http.Response, error) {
	if c.debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.client.Do(request)
	if err != nil {
		return resp, err
	}

	if c.debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}
