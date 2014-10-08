package hipchat

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.hipchat.com/v2/"
)

type Client struct {
	AuthToken string
	BaseURL   *url.URL
	client    *http.Client
}

func NewClient(authToken string) *Client {
	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		panic(err)
	}

	return &Client{
		AuthToken: authToken,
		BaseURL:   baseURL,
		client:    http.DefaultClient,
	}
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

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

	req.Header.Add("Authorization", "Bearer "+c.AuthToken)
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return resp, errors.New("Server returns status!=200")
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
