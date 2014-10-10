package main

import (
	"flag"
	"strconv"

	"github.com/tbruyelle/hipchat"
)

var (
	token  = flag.String("token", "", "The HipChat AuthToken")
	roomId = flag.String("room", "", "The HipChat room id")
)

func main() {
	flag.Parse()
	if *token == "" || *roomId == "" {
		flag.PrintDefaults()
		return
	}
	c := hipchat.NewClient(*token)

	rooms, _, err := c.Room.List()
	if err != nil {
		panic(err)
	}

	notifRq := &hipchat.NotificationRequest{Message: "Hey there!"}

	for _, room := range rooms.Items {
		_, err := c.Room.Notification(strconv.Itoa(room.ID), notifRq)
		if err != nil {
			panic(err)
		}
	}
}
