package storage

import (
	"context"
	"strider-backend-test.com/adapter/mysql"
	"strider-backend-test.com/api/routes/user/model"
)

type Repository interface {
	mysql.Repository
	IncrementPostCount(ctx context.Context, tx mysql.Transaction, request *model.UpdateRequest) error
	IncrementFollowerCount(ctx context.Context, tx mysql.Transaction, request *model.UpdateRequest) error
	IncrementFollowingCount(ctx context.Context, tx mysql.Transaction, request *model.UpdateRequest) error
}
