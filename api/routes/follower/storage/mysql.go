package storage

import (
	"context"
	"strider-backend-test.com/adapter/mysql"
	"strider-backend-test.com/api/routes/follower/model"
)

type storage struct {
	mysql.Repository
}

func New(repository mysql.Repository) Repository {
	return &storage{repository}
}

func (s *storage) Create(ctx context.Context, tx mysql.Transaction, follower *model.Follower) error {
	followerDB := model.NewFollowerDB(follower)

	insert := `
	INSERT INTO strider.follower(
			id,
			follow,
			created_at
		)
	VALUES(?, ?, ?)`

	_, err := tx.ExecContext(
		ctx,
		insert,
		followerDB.ID,
		followerDB.Follow,
		followerDB.CreatedAt)

	return err
}

func (s *storage) Delete(ctx context.Context, tx mysql.Transaction, follower *model.Follower) error {
	insert := `
	DELETE FROM strider.follower 
	WHERE id = ? AND follow = ?`

	_, err := tx.ExecContext(
		ctx,
		insert,
		follower.ID,
		follower.Follow)

	return err
}
