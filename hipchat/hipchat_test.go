package hipchat

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

// setup sets up a test HTTP server and a hipchat.Client configured to talk
// to that test server.
// Tests should register handlers on mux which provide mock responses for
// the API method being tested.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// github client configured to use test server
	client = NewClient("AuthToken")
	url, _ := url.Parse(server.URL)
	client.baseURL = url
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	authToken := "AuthToken"

	c := NewClient(authToken)

	if c.authToken != authToken {
		t.Errorf("NewClient authToken %s, want %s", c.authToken, authToken)
	}
	if c.baseURL.String() != defaultBaseURL {
		t.Error("NewClient baseURL %s, want %s", c.baseURL.String(), defaultBaseURL)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient("AuthToken")

	inURL, outURL := "foo", defaultBaseURL+"foo"
	inBody, outBody := &NotificationRequest{Message: "Hello"}, `{"message":"Hello"}`+"\n"
	r, _ := c.NewRequest("GET", inURL, inBody)

	if r.URL.String() != outURL {
		t.Errorf("NewRequest URL %s, want %s", r.URL.String(), outURL)
	}
	body, _ := ioutil.ReadAll(r.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest body %s, want %s", body, outBody)
	}
	authorization := r.Header.Get("Authorization")
	if authorization != "Bearer "+c.authToken {
		t.Errorf("NewRequest authorization header %s, want %s", authorization, "Bearer "+c.authToken)
	}
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("NewRequest Content-Type header %s, want application/json", contentType)
	}
}
