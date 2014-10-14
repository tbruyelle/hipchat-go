# HipChat

Go client library for the [HipChat API v2](https://www.hipchat.com/docs/apiv2).

[GoDoc](https://godoc.org/github.com/tbruyelle/hipchat)

Currently only a small part of the API is implemented, so pull requests are welcome.

### Usage

```go
import "github.com/tbruyelle/hipchat-go/hipchat"
```

Spam all the rooms you have access to (not recommanded):

```go
c := hipchat.NewClient("<your AuthToken here>")

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
```


---
The code architecture is hugely inspired by [google/go-github](github.com/google/go-github).


