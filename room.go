package hipchat

import (
	"fmt"
	"net/http"
)

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

// Get all rooms
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/get_all_rooms
func (c *Client) Rooms() (*Rooms, *http.Response, error) {
	req, err := c.NewRequest("GET", "room", nil)
	if err != nil {
		return nil, nil, err
	}

	rooms := new(Rooms)
	resp, err := c.Do(req, rooms)
	if err != nil {
		return nil, resp, err
	}
	return rooms, resp, nil
}

// Get room
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/get_room
func (c *Client) Room(id string) (*Room, *http.Response, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("room/%s", id), nil)
	if err != nil {
		return nil, nil, err
	}

	room := new(Room)
	resp, err := c.Do(req, room)
	if err != nil {
		return nil, resp, err
	}
	return room, resp, nil
}
