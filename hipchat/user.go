package hipchat

import (
	"fmt"
	"net/http"
)

// User represents the HipChat user.
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MentionName string `json:"mention_name"`
	Links       Links  `json:"links"`
}

// UserService gives access to the user related methods of the API.
type UserService struct {
	client *Client
}

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

