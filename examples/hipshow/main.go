package main

import (
	"flag"
	"fmt"

	"github.com/tbruyelle/hipchat-go/hipchat"
)

var (
	token = flag.String("token", "", "The HipChat AuthToken")
)

func main() {
	flag.Parse()
	if *token == "" {
		flag.PrintDefaults()
		return
	}
	c := hipchat.NewClient(*token)
	startIndex := 0
	maxResults := 5
	var allRooms []hipchat.Room

	for {
		opt := &hipchat.RoomsListOptions{
			ListOptions:     hipchat.ListOptions{StartIndex: startIndex, MaxResults: maxResults},
			IncludePrivate:  true,
			IncludeArchived: true}

		rooms, resp, err := c.Room.List(opt)

		if err != nil {
			fmt.Printf("Error during room list req %q\n", err)
			fmt.Printf("Server returns %+v\n", resp)
			return
		}

		allRooms = append(allRooms, rooms.Items...)
		if rooms.Links.Next != "" {
			startIndex += maxResults
		} else {
			break
		}
	}

	fmt.Printf("Your group has %d rooms:\n", len(allRooms))
	for _, r := range allRooms {
		fmt.Printf("%d %s \n", r.ID, r.Name)
	}
}
