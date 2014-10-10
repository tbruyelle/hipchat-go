package main

import (
	"flag"
	"fmt"

	"github.com/tbruyelle/hipchat"
)

var (
	token = flag.String("token", "", "The HipChat AuthToken")
)

func main() {
	flag.Parse()
	c := hipchat.NewClient(*token)
	fmt.Printf("%s - %+v\n", *token, c)

	rooms, resp, err := c.Room.List()
	fmt.Println(resp, err)
	fmt.Printf("\n%+v\n", rooms)
	room, resp, err := c.Room.Get("763227")
	fmt.Println(resp, err)
	fmt.Printf("\n%+v\n", room)

	n := &hipchat.NotificationRequest{Message: "(lol)", MessageFormat: "text"}
	resp, err = c.Room.Notification("763227", n)
	fmt.Println(resp, err)
}
