package follower

import (
	"context"
	"strider-backend-test.com/adapter/mysql"
	"strider-backend-test.com/api/routes/follower/model"
	"strider-backend-test.com/api/routes/follower/storage"
	"strider-backend-test.com/api/routes/user"
	usermodel "strider-backend-test.com/api/routes/user/model"
)

type Service interface {
	Follow(ctx context.Context, follower *model.Follower) (err error)
	Unfollow(ctx context.Context, follower *model.Follower) (err error)
}

type service struct {
	user.Service
	storage storage.Repository
}

func NewService(userService user.Service, storage storage.Repository) Service {
	return &service{userService, storage}
}

func (s *service) Follow(ctx context.Context, follower *model.Follower) (err error) {

	tx, err := s.storage.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			s.storage.Rollback(ctx, tx)
		} else {
			s.storage.Commit(ctx, tx)
		}
	}()

	if err := s.updateFollowingCount(ctx, tx, follower.ID, 1); err != nil {
		return err
	}

	if err := s.updateFollowerCount(ctx, tx, follower.Follow, 1); err != nil {
		return err
	}

	return s.storage.Create(ctx, tx, follower)
}

func (s *service) Unfollow(ctx context.Context, follower *model.Follower) (err error) {

	tx, err := s.storage.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			s.storage.Rollback(ctx, tx)
		} else {
			s.storage.Commit(ctx, tx)
		}
	}()

	if err := s.updateFollowingCount(ctx, tx, follower.ID, -1); err != nil {
		return err
	}

	if err := s.updateFollowerCount(ctx, tx, follower.Follow, -1); err != nil {
		return err
	}

	return s.storage.Delete(ctx, tx, follower)
}

func (s *service) updateFollowerCount(ctx context.Context, tx mysql.Transaction, userID string, count int) error {
	updateRequest := usermodel.UpdateRequest{UserID: userID, FollowerCount: count}

	return s.UpdateCount(ctx, tx, &updateRequest)
}

func (s *service) updateFollowingCount(ctx context.Context, tx mysql.Transaction, userID string, count int) error {
	updateRequest := usermodel.UpdateRequest{UserID: userID, FollowingCount: count}

	return s.UpdateCount(ctx, tx, &updateRequest)
}
