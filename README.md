# hipchat-go 

Go client library for the [HipChat API v2](https://www.hipchat.com/docs/apiv2).

[![GoDoc](https://godoc.org/github.com/tbruyelle/hipchat-go/hipchat?status.svg)](https://godoc.org/github.com/tbruyelle/hipchat-go/hipchat)
[![Build Status](https://travis-ci.org/tbruyelle/hipchat-go.svg??branch=master)](https://travis-ci.org/tbruyelle/hipchat-go)

Currently only a small part of the API is implemented, so pull requests are welcome.

### Usage

```go
import "github.com/tbruyelle/hipchat-go/hipchat"
```

Build a new client, then use the `client.Room` service to spam all the rooms you have access to (not recommended):

```go
c := hipchat.NewClient("<your AuthToken here>")

rooms, _, err := c.Room.List()
if err != nil {
	panic(err)
}

notifRq := &hipchat.NotificationRequest{Message: "Hey there!"}

for _, room := range rooms.Items {
	_, err := c.Room.Notification(room.Name, notifRq)
	if err != nil {
		panic(err)
	}
}
```

The above example creates a client that connects to the primary HipChat API served by `api.hipchat.com`. If you wish to communicate with a custom HipChat installation you may need to provide a different URL. For example a company may purchase licenses to install HipChat on corporate servers behind a firewall. In this case the URL endpoint is not at the default location. To accommodate this the base URL the `Client` structure uses may be overwritten. Below is an example of doing this. It is the same example from above, except that the URL is modified.

```go
c := hipchat.NewClient("<your AuthToken here>")
c.BaseURL = "https://hipchat.mycustomdomain.com/v2/"

rooms, _, err := c.Room.List()
if err != nil {
    panic(err)
}

notifRq := &hipchat.NotificationRequest{Message: "Hey there!"}

for _, room := range rooms.Items {
    _, err := c.Room.Notification(room.Name, notifRq)
    if err != nil {
        panic(err)
    }
}
```

---
The code architecture is hugely inspired by [google/go-github](github.com/google/go-github).


