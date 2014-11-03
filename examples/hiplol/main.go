package main

import (
	"flag"
	"fmt"

	"github.com/tbruyelle/hipchat-go/hipchat"
)

var (
	token  = flag.String("token", "", "The HipChat AuthToken")
	roomId = flag.Int("room", 0, "The HipChat room id")
)

func main() {
	flag.Parse()
	if *token == "" || *roomId == 0 {
		flag.PrintDefaults()
		return
	}
	c := hipchat.NewClient(*token)

	notifRq := &hipchat.NotificationRequest{Message: "Hey there!"}
	resp, err := c.Room.Notification(*roomId, notifRq)
	if err != nil {
		fmt.Printf("Error during room notification %q\n", err)
		fmt.Printf("Server returns %+v\n", resp)
		return
	}
	fmt.Println("Lol sent !")
}
