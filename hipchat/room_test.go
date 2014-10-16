package hipchat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRoomGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/room/1", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		fmt.Fprintf(w, `{"id":1, "name":"n", "links":{"self":"s"}}`)
	})
	want := &Room{ID: 1, Name: "n", Links: RoomLinks{Self: "s"}}

	room, _, err := client.Room.Get(1)
	if err != nil {
		t.Fatalf("Room.Get returns an error %v", err)
	}
	if !reflect.DeepEqual(want, room) {
		t.Errorf("Room.Get returned %+v, want %+v", room, want)
	}
}

func TestRoomList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method %s, want %s", r.Method, m)
		}
		fmt.Fprintf(w, `
		{
			"items": [{"id":1,"name":"n"}], 
			"startIndex":1,
			"maxResults":1,
			"links":{"Self":"s"}
		}`)
	})
	want := &Rooms{Items: []Room{Room{ID: 1, Name: "n"}}, StartIndex: 1, MaxResults: 1, Links: RoomsLinks{Self: "s"}}

	rooms, _, err := client.Room.List()
	if err != nil {
		t.Fatalf("Room.List returns an error %v", err)
	}
	if !reflect.DeepEqual(want, rooms) {
		t.Errorf("Room.List returned %+v, want %+v", rooms, want)
	}
}

func TestRoomNotification(t *testing.T) {
	setup()
	defer teardown()

	args := &NotificationRequest{Message: "m", MessageFormat: "text"}

	mux.HandleFunc("/room/1/notification", func(w http.ResponseWriter, r *http.Request) {
		if m := "POST"; m != r.Method {
			t.Errorf("Request method %s, want %s", r.Method, m)
		}
		v := new(NotificationRequest)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, args) {
			t.Errorf("Request body %+v, want %+v", v, args)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Room.Notification(1, args)
	if err != nil {
		t.Fatalf("Room.Notification returns an error %v", err)
	}
}

func TestRoomCreate(t *testing.T) {
	setup()
	defer teardown()

	args := &CreateRoomRequest{Name: "n", Topic: "t"}

	mux.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		if m := "POST"; m != r.Method {
			t.Errorf("Request method %s, want %s", r.Method, m)
		}
		v := new(CreateRoomRequest)
		json.NewDecoder(r.Body).Decode(v)

		if !reflect.DeepEqual(v, args) {
			t.Errorf("Request body %+v, want %+v", v, args)
		}
		fmt.Fprintf(w, `{"id":1,"links":{"self":"s"}}`)
	})
	want := &Room{ID: 1, Links: RoomLinks{Self: "s"}}

	room, _, err := client.Room.Create(args)
	if err != nil {
		t.Fatalf("Room.Create returns an error %v", err)
	}
	if !reflect.DeepEqual(room, want) {
		t.Errorf("Room.Create returns %+v, want %+v", room, want)
	}
}
