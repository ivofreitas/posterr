package storage

import (
	"context"
	"strider-backend-test.com/adapter/mysql"
	"strider-backend-test.com/api/routes/post/model"
)

type Repository interface {
	mysql.Repository
	Create(ctx context.Context, tx mysql.Transaction, post *model.Post) error
	List(ctx context.Context, limit, offset int) ([]*model.Post, error)
	ListByFollower(ctx context.Context, follower string, limit, offset int) ([]*model.Post, error)
	ListByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Post, error)
	GetPostCountToday(ctx context.Context, userID string) (int, error)
}
