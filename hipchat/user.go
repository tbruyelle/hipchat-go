package hipchat

import (
	"fmt"
	"net/http"
)

// UserPresence represents the HipChat user's presence.
type UserPresence struct {
	Status   string `json:"status"`
	Idle     int    `json:"idle"`
	Show     string `json:"show"`
	IsOnline bool   `json:"is_online"`
}

// User represents the HipChat user.
type User struct {
	XmppJid      string       `json:"xmpp_jid"`
	IsDeleted    bool         `json:"is_deleted"`
	Name         string       `json:"name"`
	LastActive   string       `json:"last_active"`
	Title        string       `json:"title"`
	Presence     UserPresence `json:"presence"`
	Created      string       `json:"created"`
	ID           int          `json:"id"`
	MentionName  string       `json:"mention_name"`
	IsGroupAdmin bool         `json:"is_group_admin"`
	Timezone     string       `json:"timezone"`
	IsGuest      bool         `json:"is_guest"`
	Email        string       `json:"email"`
	PhotoUrl     string       `json:"photo_url"`
	Links        Links        `json:"links"`
}

// UserService gives access to the user related methods of the API.
type UserService struct {
	client *Client
}

// UserRequest represents a HipChat user request.
type UserRequest struct{}

// ShareFile sends a file to the user specified by the id.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/share_file_with_user
func (u *UserService) ShareFile(id string, shareFileReq *ShareFileRequest) (*http.Response, error) {
	req, err := u.client.NewFileUploadRequest("POST", fmt.Sprintf("user/%s/share/file", id), shareFileReq)
	if err != nil {
		return nil, err
	}

	return u.client.Do(req, nil)
}

// View fetches a user's details.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/view_user
func (u *UserService) View(id string, userReq *UserRequest) (*User, *http.Response, error) {
	req, err := u.client.NewRequest("GET", fmt.Sprintf("user/%s", id), userReq)

	userDetails := new(User)
	resp, err := u.client.Do(req, &userDetails)
	if err != nil {
		return nil, resp, err
	}
	return userDetails, resp, nil
}
