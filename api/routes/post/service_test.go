package post

import (
	"errors"
	"github.com/golang/mock/gomock"
	mysqlmock "strider-backend-test.com/adapter/mysql/mock"
	"strider-backend-test.com/api/routes/post/mock"
	usermock "strider-backend-test.com/api/routes/user/mock"
	"testing"
)

func TestSvCreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	gpct := repository.EXPECT().GetPostCountToday(mock.Ctx, mock.PostRequest.CreatedBy).After(bt).Return(2, nil)
	uc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementPostCount).After(gpct).Return(nil)
	c := repository.EXPECT().Create(mock.Ctx, tx, &mock.Post).After(uc).Return(nil)
	repository.EXPECT().Commit(mock.Ctx, tx).After(c)

	service := NewService(userService, repository)
	if _, err := service.Create(mock.Ctx, &mock.PostRequest); err != nil {
		t.Fail()
	}
}

func TestSvCreateQuote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	gpct := repository.EXPECT().GetPostCountToday(mock.Ctx, mock.QuoteRequest.CreatedBy).After(bt).Return(2, nil)
	uc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementPostCount).After(gpct).Return(nil)
	c := repository.EXPECT().Create(mock.Ctx, tx, &mock.Quote).After(uc).Return(nil)
	repository.EXPECT().Commit(mock.Ctx, tx).After(c)

	service := NewService(userService, repository)
	if _, err := service.Create(mock.Ctx, &mock.QuoteRequest); err != nil {
		t.Fail()
	}
}

func TestSvCreateRepost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	gpct := repository.EXPECT().GetPostCountToday(mock.Ctx, mock.RepostRequest.CreatedBy).After(bt).Return(2, nil)
	uc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementPostCount).After(gpct).Return(nil)
	c := repository.EXPECT().Create(mock.Ctx, tx, &mock.Repost).After(uc).Return(nil)
	repository.EXPECT().Commit(mock.Ctx, tx).After(c)

	service := NewService(userService, repository)
	if _, err := service.Create(mock.Ctx, &mock.RepostRequest); err != nil {
		t.Fail()
	}
}

func TestSvCreateBTError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)

	btErr := errors.New("generic error")
	repository.EXPECT().BeginTransaction(mock.Ctx).Return(nil, btErr)

	service := NewService(userService, repository)
	if _, err := service.Create(mock.Ctx, &mock.PostRequest); err != btErr {
		t.Fail()
	}
}

func TestSvCreateGPCTError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	gpctErr := errors.New("generic error")
	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	gpct := repository.EXPECT().GetPostCountToday(mock.Ctx, mock.PostRequest.CreatedBy).After(bt).Return(0, gpctErr)
	repository.EXPECT().Rollback(mock.Ctx, tx).After(gpct)

	service := NewService(userService, repository)
	if _, err := service.Create(mock.Ctx, &mock.PostRequest); err != gpctErr {
		t.Fail()
	}
}

func TestSvCreatePostNotAllowed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	fivePostErr := errors.New("not allowed to post more than 5 times")
	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	gpct := repository.EXPECT().GetPostCountToday(mock.Ctx, mock.PostRequest.CreatedBy).After(bt).Return(5, nil)
	repository.EXPECT().Rollback(mock.Ctx, tx).After(gpct)

	service := NewService(userService, repository)
	if _, err := service.Create(mock.Ctx, &mock.PostRequest); err.Error() != fivePostErr.Error() {
		t.Fail()
	}
}

func TestSvCreateUCError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	ucErr := errors.New("generic error")
	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	gpct := repository.EXPECT().GetPostCountToday(mock.Ctx, mock.PostRequest.CreatedBy).After(bt).Return(1, nil)
	uc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementPostCount).After(gpct).Return(ucErr)
	repository.EXPECT().Rollback(mock.Ctx, tx).After(uc)

	service := NewService(userService, repository)
	if _, err := service.Create(mock.Ctx, &mock.PostRequest); err != ucErr {
		t.Fail()
	}
}

func TestSvCreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)
	tx := mysqlmock.NewMockTransaction(ctrl)

	ucErr := errors.New("generic error")
	bt := repository.EXPECT().BeginTransaction(mock.Ctx).Return(tx, nil)
	gpct := repository.EXPECT().GetPostCountToday(mock.Ctx, mock.PostRequest.CreatedBy).After(bt).Return(1, nil)
	uc := userService.EXPECT().UpdateCount(mock.Ctx, tx, &usermock.IncrementPostCount).After(gpct).Return(nil)
	c := repository.EXPECT().Create(mock.Ctx, tx, &mock.Post).After(uc).Return(ucErr)
	repository.EXPECT().Rollback(mock.Ctx, tx).After(c)

	service := NewService(userService, repository)
	if _, err := service.Create(mock.Ctx, &mock.PostRequest); err != ucErr {
		t.Fail()
	}
}

func TestSvListUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)

	repository.EXPECT().ListByUser(mock.Ctx, mock.ListUserRequest.User, mock.ListUserRequest.Limit, mock.ListUserRequest.Offset).Return(mock.Posts, nil)

	service := NewService(userService, repository)
	if _, err := service.List(mock.Ctx, &mock.ListUserRequest); err != nil {
		t.Fail()
	}
}

func TestSvListFollower(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)

	repository.EXPECT().ListByFollower(mock.Ctx, mock.ListFollowerRequest.Follower, mock.ListFollowerRequest.Limit, mock.ListFollowerRequest.Offset).Return(mock.Posts, nil)

	service := NewService(userService, repository)
	if _, err := service.List(mock.Ctx, &mock.ListFollowerRequest); err != nil {
		t.Fail()
	}
}

func TestSvListAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)

	repository.EXPECT().List(mock.Ctx, mock.ListAllRequest.Limit, mock.ListAllRequest.Offset).Return(mock.Posts, nil)

	service := NewService(userService, repository)
	if _, err := service.List(mock.Ctx, &mock.ListAllRequest); err != nil {
		t.Fail()
	}
}

func TestSvListAllError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := usermock.NewMockService(ctrl)
	repository := mock.NewMockRepository(ctrl)

	listErr := errors.New("generic error")
	repository.EXPECT().List(mock.Ctx, mock.ListAllRequest.Limit, mock.ListAllRequest.Offset).Return(nil, listErr)

	service := NewService(userService, repository)
	if _, err := service.List(mock.Ctx, &mock.ListAllRequest); err != listErr {
		t.Fail()
	}
}
