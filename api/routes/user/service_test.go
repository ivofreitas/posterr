package user

import (
	"errors"
	"github.com/golang/mock/gomock"
	mysqlmock "strider-backend-test.com/adapter/mysql/mock"
	"strider-backend-test.com/api/routes/user/mock"
	"testing"
)

func TestSvUpdatePostCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	repository.EXPECT().IncrementPostCount(mock.Ctx, tx, &mock.IncrementPostCount).Return(nil)

	service := NewService(repository)
	if err := service.UpdateCount(mock.Ctx, tx, &mock.IncrementPostCount); err != nil {
		t.Fail()
	}
}

func TestSvUpdatePostCountError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	postCountErr := errors.New("post count error")
	repository.EXPECT().IncrementPostCount(mock.Ctx, tx, &mock.IncrementPostCount).Return(postCountErr)

	service := NewService(repository)
	if err := service.UpdateCount(mock.Ctx, tx, &mock.IncrementPostCount); err.Error() != postCountErr.Error() {
		t.Fail()
	}
}

func TestSvUpdateFollowerCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	repository.EXPECT().IncrementFollowerCount(mock.Ctx, tx, &mock.IncrementFollowerCount).Return(nil)

	service := NewService(repository)
	if err := service.UpdateCount(mock.Ctx, tx, &mock.IncrementFollowerCount); err != nil {
		t.Fail()
	}
}

func TestSvUpdateFollowerCountError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	followerCountErr := errors.New("follower count error")
	repository.EXPECT().IncrementFollowerCount(mock.Ctx, tx, &mock.IncrementFollowerCount).Return(followerCountErr)

	service := NewService(repository)
	if err := service.UpdateCount(mock.Ctx, tx, &mock.IncrementFollowerCount); err.Error() != followerCountErr.Error() {
		t.Fail()
	}
}

func TestSvUpdateFollowingCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	repository.EXPECT().IncrementFollowingCount(mock.Ctx, tx, &mock.IncrementFollowingCount).Return(nil)

	service := NewService(repository)
	if err := service.UpdateCount(mock.Ctx, tx, &mock.IncrementFollowingCount); err != nil {
		t.Fail()
	}
}

func TestSvUpdateFollowingCountError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	followingCountErr := errors.New("following count error")
	repository.EXPECT().IncrementFollowingCount(mock.Ctx, tx, &mock.IncrementFollowingCount).Return(followingCountErr)

	service := NewService(repository)
	if err := service.UpdateCount(mock.Ctx, tx, &mock.IncrementFollowingCount); err.Error() != followingCountErr.Error() {
		t.Fail()
	}
}
