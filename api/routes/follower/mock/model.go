package mock

import (
	"context"
	"strider-backend-test.com/api/routes/follower/model"
)

var (
	Request = model.Request{
		Follower: "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		Follow:   "sad357981-35sd-ds12-sd12-4asd654sd78",
	}
	BrokenRequest     = `{"following":123}`
	RequestIncomplete = model.Request{
		Follower: "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
	}
	Follower = model.Follower{
		ID:     "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		Follow: "sad357981-35sd-ds12-sd12-4asd654sd78"}
	Ctx = context.Background()
)
