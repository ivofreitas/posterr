package mock

import (
	"context"
	"strider-backend-test.com/api/routes/post/model"
)

var (
	Post = model.Post{
		Content:   "Content 1, 2, 3...",
		CreatedBy: "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23"}
	Posts = []*model.Post{&Post}
	Quote = model.Post{
		Content:   "Content 1, 2, 3...",
		CreatedBy: "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		Parent:    &model.Post{ID: "sad357981-35sd-ds12-sd12-4asd654sd78"}}
	Repost = model.Post{
		CreatedBy: "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		Parent:    &model.Post{ID: "sad357981-35sd-ds12-sd12-4asd654sd78"}}
	EmptyPost   *model.Post
	EmptyPosts  []*model.Post
	Ctx         = context.Background()
	PostRequest = model.Request{
		Content:   "Content 1, 2, 3...",
		CreatedBy: "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23"}
	PostRequestWithoutCreator = model.Request{
		Content: "Content 1, 2, 3..."}
	BrokenRequest = `{"content":123}`
	QuoteRequest  = model.Request{
		Content:   "Content 1, 2, 3...",
		CreatedBy: "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		ParentID:  "sad357981-35sd-ds12-sd12-4asd654sd78"}
	RepostRequest = model.Request{
		CreatedBy: "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		ParentID:  "sad357981-35sd-ds12-sd12-4asd654sd78"}
	EmptyRequest    *model.Request
	ListUserRequest = model.ListRequest{
		User:   "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		Offset: 0,
		Limit:  50,
	}
	ListFollowerRequest = model.ListRequest{
		Follower: "dc5b9b53-0bb1-45d0-9eac-f441dcc16d23",
		Offset:   0,
		Limit:    50,
	}
	ListAllRequest = model.ListRequest{
		Offset: 0,
		Limit:  50,
	}
)
