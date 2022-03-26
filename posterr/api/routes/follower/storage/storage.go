package storage

import (
	"context"
	"strider-backend-test.com/adapter/mysql"
	"strider-backend-test.com/api/routes/follower/model"
)

type Repository interface {
	mysql.Repository
	Create(ctx context.Context, tx mysql.Transaction, follower *model.Follower) error
	Delete(ctx context.Context, tx mysql.Transaction, follower *model.Follower) error
}
