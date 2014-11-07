package hipchat

import (
	"fmt"
	"net/http"
)

// RoomService gives access to the room related methods of the API.
type RoomService struct {
	client *Client
}

// Rooms represents a HipChat room list.
type Rooms struct {
	Items      []Room    `json:"items"`
	StartIndex int       `json:"startIndex"`
	MaxResults int       `json:"maxResults"`
	Links      PageLinks `json:"links"`
}

// Room represents a HipChat room.
type Room struct {
	ID                int            `json:"id"`
	Links             RoomLinks      `json:"links"`
	Name              string         `json:"name"`
	XmppJid           string         `json:"xmpp_jid"`
	Statistics        RoomStatistics `json:"statistics"`
	Created           string         `json:"created"`
	IsArchived        bool           `json:"is_archived"`
	Privacy           string         `json:"privacy"`
	IsGuestAccessible bool           `json:"is_guess_accessible"`
	Topic             string         `json:"topic"`
	Participants      []User         `json:"participants"`
	Owner             User           `json:"owner"`
	GuestAccessURL    string         `json:"guest_access_url"`
}

// RoomStatistics represents the HipChat room statistics.
type RoomStatistics struct {
	Links Links `json:"links"`
}

// CreateRoomRequest represents a HipChat room creation request.
type CreateRoomRequest struct {
	Topic       string `json:"topic,omitempty"`
	GuestAccess bool   `json:"guest_access,omitempty"`
	Name        string `json:"name,omitempty"`
	OwnerUserID string `json:"owner_user_id,omitempty"`
	Privacy     string `json:"privacy,omitempty"`
}

// UpdateRoomRequest represents a HipChat room update request.
type UpdateRoomRequest struct {
	Name          string `json:"name"`
	Topic         string `json:"topic"`
	IsGuestAccess bool   `json:"is_guest_access"`
	IsArchived    bool   `json:"is_archived"`
	Privacy       string `json:"privacy"`
	Owner         ID     `json:"owner"`
}

// RoomLinks represents the HipChat room links.
type RoomLinks struct {
	Links
	Webhooks     string `json:"webhooks"`
	Members      string `json:"members"`
	Participants string `json:"participants"`
}

// NotificationRequest represents a HipChat room notification request.
type NotificationRequest struct {
	Color         string `json:"color,omitempty"`
	Message       string `json:"message,omitempty"`
	Notify        bool   `json:"notify,omitempty"`
	MessageFormat string `json:"message_format,omitempty"`
}

// List returns all the rooms authorized.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/get_all_rooms
func (r *RoomService) List() (*Rooms, *http.Response, error) {
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

// Get returns the room specified by the id.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/get_room
func (r *RoomService) Get(id string) (*Room, *http.Response, error) {
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

// Notification sends a notification to the room specified by the id.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/send_room_notification
func (r *RoomService) Notification(id string, notifReq *NotificationRequest) (*http.Response, error) {
	req, err := r.client.NewRequest("POST", fmt.Sprintf("room/%s/notification", id), notifReq)
	if err != nil {
		return nil, err
	}

	return r.client.Do(req, nil)
}

// Create creates a new room.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/create_room
func (r *RoomService) Create(roomReq *CreateRoomRequest) (*Room, *http.Response, error) {
	req, err := r.client.NewRequest("POST", "room", roomReq)
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

// Update updates an existing room.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/update_room
func (r *RoomService) Update(id string, roomReq *UpdateRoomRequest) (*http.Response, error) {
	req, err := r.client.NewRequest("PUT", fmt.Sprintf("room/%s", id), roomReq)
	if err != nil {
		return nil, err
	}

	return r.client.Do(req, nil)
}

// List all webhooks for a given room.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/get_all_webhooks
func (r *RoomService) GetAllWebhooks(id interface{}, roomReq *GetAllWebhooksRequest) (*WebhookList, *http.Response, error) {
	req, err := r.client.NewRequest("GET", fmt.Sprintf("room/%v/webhook", id), roomReq)
	if err != nil {
		return nil, nil, err
	}

	whList := new(WebhookList)

	resp, err := r.client.Do(req, whList)
	if err != nil {
		return nil, resp, err
	}
	return whList, resp, nil
}

// Delete a given webhook.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/delete_webhook
func (r *RoomService) DeleteWebhook(id interface{}, webhookId interface{}) (*http.Response, error) {
	req, err := r.client.NewRequest("DELETE", fmt.Sprintf("room/%v/webhook/%v", id, webhookId), nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Create a new webhook.
//
// HipChat API docs: https://www.hipchat.com/docs/apiv2/method/create_webhook
func (r *RoomService) CreateWebhook(id interface{}, roomReq *CreateWebhookRequest) (*Webhook, *http.Response, error) {
	req, err := r.client.NewRequest("POST", fmt.Sprintf("room/%v/webhook", id), roomReq)
	if err != nil {
		return nil, nil, err
	}

	wh := new(Webhook)

	resp, err := r.client.Do(req, wh)
	if err != nil {
		return nil, resp, err
	}

	return wh, resp, nil
}
