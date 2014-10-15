package hipchat

import (
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
