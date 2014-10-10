package main

import (
	"flag"
	"fmt"

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

	n := &hipchat.NotificationRequest{Message: "(lol)", MessageFormat: "text"}
	resp, err := c.Room.Notification(*roomId, n)
	if err != nil {
		fmt.Printf("Error during room notification %q\n", err)
		fmt.Printf("Server returns %+v\n", resp)
		return
	}
	fmt.Println("lol sent !")
}
