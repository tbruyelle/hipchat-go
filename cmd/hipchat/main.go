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

	v, resp, err := c.Rooms()
	fmt.Println(resp, err)
	fmt.Printf("\n%+v\n", v)

}
