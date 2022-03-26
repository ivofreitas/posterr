package post

import (
	"context"
	"errors"
	"strider-backend-test.com/adapter/mysql"
	"strider-backend-test.com/api/routes/post/model"
	"strider-backend-test.com/api/routes/post/storage"
	"strider-backend-test.com/api/routes/user"
	usermodel "strider-backend-test.com/api/routes/user/model"
)

type Service interface {
	Create(ctx context.Context, request *model.Request) (post *model.Post, err error)
	List(ctx context.Context, listRequest *model.ListRequest) (posts []*model.Post, err error)
}

type service struct {
	user.Service
	storage storage.Repository
}

func NewService(userService user.Service, storage storage.Repository) Service {
	return &service{userService, storage}
}

func (s *service) Create(ctx context.Context, request *model.Request) (post *model.Post, err error) {

	tx, err := s.storage.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			s.storage.Rollback(ctx, tx)
		} else {
			s.storage.Commit(ctx, tx)
		}
	}()

	err = s.updateUserCount(ctx, tx, request)
	if err != nil {
		return nil, err
	}

	post = new(model.Post)
	post.CreatedBy = request.CreatedBy
	if len(request.Content) > 0 {
		post.Content = request.Content
	}
	if len(request.ParentID) > 0 {
		post.Parent = new(model.Post)
		post.Parent.ID = request.ParentID
	}

	if err := s.storage.Create(ctx, tx, post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *service) updateUserCount(ctx context.Context, tx mysql.Transaction, request *model.Request) (err error) {
	postsCount, err := s.storage.GetPostCountToday(ctx, request.CreatedBy)
	if err != nil {
		return err
	}

	if postsCount >= 5 {
		return errors.New("not allowed to post more than 5 times")
	}

	updateRequest := usermodel.UpdateRequest{UserID: request.CreatedBy, PostsCount: 1}

	return s.UpdateCount(ctx, tx, &updateRequest)
}

func (s *service) List(ctx context.Context, listRequest *model.ListRequest) (posts []*model.Post, err error) {

	if len(listRequest.User) > 0 {
		posts, err = s.storage.ListByUser(ctx, listRequest.User, listRequest.Limit, listRequest.Offset)
	} else if (len(listRequest.Follower)) > 0 {
		posts, err = s.storage.ListByFollower(ctx, listRequest.Follower, listRequest.Limit, listRequest.Offset)
	} else {
		posts, err = s.storage.List(ctx, listRequest.Limit, listRequest.Offset)
	}

	if err != nil {
		return nil, err
	}

	return posts, nil
}
