package hipchat

import (
	"fmt"
	"net/http"
)

type roomService struct {
	client *Client
}

type Rooms struct {
	Items      []Room     `json:"items"`
	StartIndex int        `json:"startIndex"`
	MaxResults int        `json:"maxResults"`
	Links      RoomsLinks `json:"links"`
}

type RoomsLinks struct {
	Self string `json:"self"`
	Prev string `json:"prev"`
	Next string `json:"next"`
}

type Room struct {
	Id    int       `json:"id"`
	Links RoomLinks `json:"links"`
	Name  string    `json:"name"`
}

type RoomLinks struct {
	Self         string `json:"self"`
	Webhooks     string `json:"webhooks"`
	Members      string `json:"members"`
	Participants string `json:"participants"`
}

type NotificationRequest struct {
	Color         *string `json:"color,omitempty"`
	Message       *string `json:"message,omitempty"`
	Notify        *bool   `json:"notify,omitempty"`
	MessageFormat *string `json:"message_format,omitempty"`
}

// Get all rooms
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/get_all_rooms
func (r *roomService) List() (*Rooms, *http.Response, error) {
	req, err := r.client.NewRequest("GET", "room", nil)
	if err != nil {
		return nil, nil, err
	}

	rooms := new(Rooms)
	resp, err := r.client.Do(req, rooms)
	if err != nil {
		return nil, resp, err
	}
	return rooms, resp, nil
}

// Get room
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/get_room
func (r *roomService) Get(id string) (*Room, *http.Response, error) {
	req, err := r.client.NewRequest("GET", fmt.Sprintf("room/%s", id), nil)
	if err != nil {
		return nil, nil, err
	}

	room := new(Room)
	resp, err := r.client.Do(req, room)
	if err != nil {
		return nil, resp, err
	}
	return room, resp, nil
}

// Send room notification
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/send_room_notification
func (r *roomService) Notification(id string, notifReq *NotificationRequest) (*http.Response, error) {
	req, err := r.client.NewRequest("POST", fmt.Sprintf("room/%s/notification", id), notifReq)
	if err != nil {
		return nil, err
	}

	return r.client.Do(req, nil)
}
