package user

import (
	"context"
	"strider-backend-test.com/adapter/mysql"
	"strider-backend-test.com/api/routes/user/model"
	"strider-backend-test.com/api/routes/user/storage"
)

type Service interface {
	UpdateCount(ctx context.Context, tx mysql.Transaction, request *model.UpdateRequest) (err error)
}

type service struct {
	storage storage.Repository
}

func NewService(storage storage.Repository) Service {
	return &service{storage}
}

func (s *service) UpdateCount(ctx context.Context, tx mysql.Transaction, request *model.UpdateRequest) (err error) {

	if request.PostsCount > 0 {
		err = s.storage.IncrementPostCount(ctx, tx, request)
	} else if request.FollowerCount > 0 {
		err = s.storage.IncrementFollowerCount(ctx, tx, request)
	} else {
		err = s.storage.IncrementFollowingCount(ctx, tx, request)
	}

	return err
}
