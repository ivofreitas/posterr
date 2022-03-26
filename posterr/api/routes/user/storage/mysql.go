package storage

import (
	"context"
	"strider-backend-test.com/adapter/mysql"
	"strider-backend-test.com/api/routes/user/model"
)

type storage struct {
	mysql.Repository
}

func New(repository mysql.Repository) Repository {
	return &storage{repository}
}

func (s *storage) IncrementPostCount(ctx context.Context, tx mysql.Transaction, request *model.UpdateRequest) error {

	update := `
	UPDATE strider.users SET posts_count = posts_count + ?
	WHERE id = ?`

	_, err := tx.ExecContext(
		ctx,
		update,
		request.PostsCount,
		request.UserID)

	return err
}

func (s *storage) IncrementFollowerCount(ctx context.Context, tx mysql.Transaction, request *model.UpdateRequest) error {

	update := `
	UPDATE strider.users SET followers_count = followers_count + ?
	WHERE id = ?`

	_, err := tx.ExecContext(
		ctx,
		update,
		request.FollowerCount,
		request.UserID)

	return err
}

func (s *storage) IncrementFollowingCount(ctx context.Context, tx mysql.Transaction, request *model.UpdateRequest) error {

	update := `
	UPDATE strider.users SET following_count = following_count + ?
	WHERE id = ?`

	_, err := tx.ExecContext(
		ctx,
		update,
		request.FollowingCount,
		request.UserID)

	return err
}
