package hipchat

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{
			"username":      "username",
			"grant_type":    "grant_type",
			"code":          "code",
			"client_name":   "client_name",
			"redirect_uri":  "redirect_uri",
			"scope":         "scope",
			"password":      "password",
			"group_id":      "group_id",
			"refresh_token": "refresh_token",
		})
		fmt.Fprintf(w, `
		{
            "access_token": "GeneratedAuthToken",
            "expires_in": 3599,
            "group_id": 123456,
            "group_name": "TestGroup",
            "scope": "send_notification view_room",
            "token_type": "bearer"
        }
        `)
	})
	want := &OAuthAccessToken{
		AccessToken: "GeneratedAuthToken",
		ExpiresIn:   3599,
		GroupID:     123456,
		GroupName:   "TestGroup",
		Scope:       "send_notification view_room",
		TokenType:   "bearer",
	}

	opt := &GenerateTokenOptions{
		"username", "grant_type", "code", "client_name", "redirect_uri",
		"scope", "password", "group_id", "refresh_token",
	}
	client.authToken = ""
	token, _, err := client.GenerateToken(opt)
	if err != nil {
		t.Fatalf("Client.GetAccessToken returns an error %v", err)
	}
	if !reflect.DeepEqual(want, token) {
		t.Errorf("Client.GetAccessToken returned %+v, want %+v", token, want)
	}
	if client.authToken != want.AccessToken {
		t.Errorf("Client.authToken = %s, want %s", client.authToken, want.AccessToken)
	}
}
