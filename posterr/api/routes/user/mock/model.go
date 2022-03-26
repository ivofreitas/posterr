package mock

import (
	"context"
	"strider-backend-test.com/api/routes/user/model"
)

var (
	IncrementPostCount = model.UpdateRequest{
		UserID:     "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		PostsCount: 1,
	}
	IncrementFollowerCount = model.UpdateRequest{
		UserID:        "sad357981-35sd-ds12-sd12-4asd654sd78",
		FollowerCount: 1,
	}
	IncrementFollowingCount = model.UpdateRequest{
		UserID:         "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		FollowingCount: 1,
	}
	DecrementPostCount = model.UpdateRequest{
		UserID:     "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		PostsCount: -1,
	}
	DecrementFollowerCount = model.UpdateRequest{
		UserID:        "sad357981-35sd-ds12-sd12-4asd654sd78",
		FollowerCount: -1,
	}
	DecrementFollowingCount = model.UpdateRequest{
		UserID:         "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		FollowingCount: -1,
	}
	Ctx = context.Background()
)
