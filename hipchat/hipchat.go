// Package hipchat provides a client for using the HipChat API v2.
package hipchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.hipchat.com/v2/"
)

// Client manages the communication with the HipChat API.
type Client struct {
	authToken string
	baseURL   *url.URL
	client    *http.Client
	// Room gives access to the /room part of the API.
	Room *RoomService
}

// NewClient returns a new HipChat API client. You must provide a valid
// AuthToken retrieved from your HipChat account.
func NewClient(authToken string) *Client {
	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		panic(err)
	}

	c := &Client{
		authToken: authToken,
		baseURL:   baseURL,
		client:    http.DefaultClient,
	}
	c.Room = &RoomService{client: c}
	return c
}

// NewRequest creates an API request. This method can be used to performs
// API request not implemented in this library. Otherwise it should not be
// be used directly.
// Relative URLs should always be specified without a preceding slash.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.authToken)
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

// Do performs the request, the json received in the response is decoded
// and stored in the value pointed by v.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if c := resp.StatusCode; c < 200 || c > 299 {
		return resp, fmt.Errorf("Server returns status %d", c)
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return resp, err
}
