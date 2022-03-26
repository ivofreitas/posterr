package follower

import (
	"errors"
	"github.com/golang/mock/gomock"
	mysqlmock "strider-backend-test.com/adapter/mysql/mock"
	"strider-backend-test.com/api/routes/follower/mock"
	usermock "strider-backend-test.com/api/routes/user/mock"
	"testing"
)

func TestSvFollow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	fiuc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementFollowingCount).After(bt).Return(nil)
	feuc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementFollowerCount).After(fiuc).Return(nil)
	c := repository.EXPECT().Create(mock.Ctx, tx, &mock.Follower).After(feuc).Return(nil)
	repository.EXPECT().Commit(mock.Ctx, tx).After(c)

	service := NewService(userService, repository)
	if err := service.Follow(mock.Ctx, &mock.Follower); err != nil {
		t.Fail()
	}
}

func TestSvFollowBTError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)

	btErr := errors.New("uc error")
	repository.EXPECT().BeginTransaction(mock.Ctx).Return(nil, btErr)

	service := NewService(userService, repository)
	if err := service.Follow(mock.Ctx, &mock.Follower); err.Error() != btErr.Error() {
		t.Fail()
	}
}

func TestSvFollowUpdateFollowingCountError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	ucErr := errors.New("uc error")
	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	uc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementFollowingCount).After(bt).Return(nil)
	feuc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementFollowerCount).After(uc).Return(ucErr)
	repository.EXPECT().Rollback(mock.Ctx, tx).After(feuc)

	service := NewService(userService, repository)
	if err := service.Follow(mock.Ctx, &mock.Follower); err.Error() != ucErr.Error() {
		t.Fail()
	}
}

func TestSvFollowUpdateFollowerCountError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	ucErr := errors.New("bt error")
	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	uc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementFollowingCount).After(bt).Return(ucErr)
	repository.EXPECT().Rollback(mock.Ctx, tx).After(uc)

	service := NewService(userService, repository)
	if err := service.Follow(mock.Ctx, &mock.Follower); err.Error() != ucErr.Error() {
		t.Fail()
	}
}

func TestSvFollowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	ucErr := errors.New("bt error")
	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	fiuc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementFollowingCount).After(bt).Return(nil)
	feuc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementFollowerCount).After(fiuc).Return(nil)
	c := repository.EXPECT().Create(mock.Ctx, tx, &mock.Follower).After(feuc).Return(ucErr)
	repository.EXPECT().Rollback(mock.Ctx, tx).After(c)

	service := NewService(userService, repository)
	if err := service.Follow(mock.Ctx, &mock.Follower); err.Error() != ucErr.Error() {
		t.Fail()
	}
}

func TestSvUnfollow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	fiuc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.DecrementFollowingCount).After(bt).Return(nil)
	feuc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.DecrementFollowerCount).After(fiuc).Return(nil)
	c := repository.EXPECT().Delete(mock.Ctx, tx, &mock.Follower).After(feuc).Return(nil)
	repository.EXPECT().Commit(mock.Ctx, tx).After(c)

	service := NewService(userService, repository)
	if err := service.Unfollow(mock.Ctx, &mock.Follower); err != nil {
		t.Fail()
	}
}

func TestSvUnfollowBTError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)

	btErr := errors.New("uc error")
	repository.EXPECT().BeginTransaction(mock.Ctx).Return(nil, btErr)

	service := NewService(userService, repository)
	if err := service.Unfollow(mock.Ctx, &mock.Follower); err.Error() != btErr.Error() {
		t.Fail()
	}
}

func TestSvUnfollowDecrementFollowingCountError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	ucErr := errors.New("uc error")
	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	uc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.DecrementFollowingCount).After(bt).Return(nil)
	feuc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.DecrementFollowerCount).After(uc).Return(ucErr)
	repository.EXPECT().Rollback(mock.Ctx, tx).After(feuc)

	service := NewService(userService, repository)
	if err := service.Unfollow(mock.Ctx, &mock.Follower); err.Error() != ucErr.Error() {
		t.Fail()
	}
}

func TestSvUnfollowDecrementFollowerCountError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	ucErr := errors.New("bt error")
	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	uc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.DecrementFollowingCount).After(bt).Return(ucErr)
	repository.EXPECT().Rollback(mock.Ctx, tx).After(uc)

	service := NewService(userService, repository)
	if err := service.Unfollow(mock.Ctx, &mock.Follower); err.Error() != ucErr.Error() {
		t.Fail()
	}
}

func TestSvUnfollowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	ucErr := errors.New("bt error")
	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	fiuc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.DecrementFollowingCount).After(bt).Return(nil)
	feuc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.DecrementFollowerCount).After(fiuc).Return(nil)
	c := repository.EXPECT().Delete(mock.Ctx, tx, &mock.Follower).After(feuc).Return(ucErr)
	repository.EXPECT().Rollback(mock.Ctx, tx).After(c)

	service := NewService(userService, repository)
	if err := service.Unfollow(mock.Ctx, &mock.Follower); err.Error() != ucErr.Error() {
		t.Fail()
	}
}
