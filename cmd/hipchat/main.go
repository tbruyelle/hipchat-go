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

	v, resp, err := c.Room("763227")
	fmt.Println(resp, err)
	fmt.Printf("\n%+v\n", v)

	n := &hipchat.NotificationRequest{Message: hipchat.String("(lol)")}
	resp, err = c.Notification("763227", n)
	fmt.Println(resp, err)
}
