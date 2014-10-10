# HipChat Go client

Go client library for the [HipChat API v2](https://www.hipchat.com/docs/apiv2).

Currently only a small part of the API is implemented, so pull requests are welcome.

### Get the library

```bash
go get github.com/tbruyelle/hipchat
```

### Usage

Spam all the rooms you have access to.

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



