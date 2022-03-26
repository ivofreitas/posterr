package follower

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strider-backend-test.com/api/middleware"
	"strider-backend-test.com/api/routes/follower/mock"
	"strings"
	"testing"
)

func TestCtlFollow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	b, _ := json.Marshal(&mock.Request)
	req := httptest.NewRequest(http.MethodPost, "/v1/strider/followers", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, ctx := prepare(req)

	service := mock.NewMockService(ctrl)
	service.EXPECT().Follow(ctx, gomock.Any()).Return(nil)
	controller := NewController(service)

	if assert.NoError(t, controller.Follow(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestCtlFollowBindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest(http.MethodPost, "/v1/strider/followers", strings.NewReader(mock.BrokenRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, _ := prepare(req)

	service := mock.NewMockService(ctrl)
	controller := NewController(service)

	if assert.NoError(t, controller.Follow(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCtlFollowValidateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	b, _ := json.Marshal(&mock.RequestIncomplete)
	req := httptest.NewRequest(http.MethodPost, "/v1/strider/followers", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, _ := prepare(req)

	service := mock.NewMockService(ctrl)
	controller := NewController(service)

	if assert.NoError(t, controller.Follow(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCtlFollowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	b, _ := json.Marshal(&mock.Request)
	req := httptest.NewRequest(http.MethodPost, "/v1/strider/followers", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, ctx := prepare(req)

	followErr := errors.New("follow error")
	service := mock.NewMockService(ctrl)
	service.EXPECT().Follow(ctx, gomock.Any()).Return(followErr)
	controller := NewController(service)

	if assert.NoError(t, controller.Follow(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}

func prepare(req *http.Request) (*httptest.ResponseRecorder, echo.Context, context.Context) {
	e := echo.New()
	e.Binder = middleware.NewBinder()
	e.Validator = middleware.NewValidator()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c, c.Request().Context()
}

func TestCtlUnfollow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	b, _ := json.Marshal(&mock.Request)
	req := httptest.NewRequest(http.MethodPost, "/v1/strider/followers", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, ctx := prepare(req)

	service := mock.NewMockService(ctrl)
	service.EXPECT().Unfollow(ctx, gomock.Any()).Return(nil)
	controller := NewController(service)

	if assert.NoError(t, controller.Unfollow(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCtlUnfollowBindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest(http.MethodPost, "/v1/strider/followers", strings.NewReader(mock.BrokenRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, _ := prepare(req)

	service := mock.NewMockService(ctrl)
	controller := NewController(service)

	if assert.NoError(t, controller.Unfollow(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCtlUnfollowValidateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	b, _ := json.Marshal(&mock.RequestIncomplete)
	req := httptest.NewRequest(http.MethodPost, "/v1/strider/followers", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, _ := prepare(req)

	service := mock.NewMockService(ctrl)
	controller := NewController(service)

	if assert.NoError(t, controller.Unfollow(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCtlUnfollowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	b, _ := json.Marshal(&mock.Request)
	req := httptest.NewRequest(http.MethodPost, "/v1/strider/followers", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, ctx := prepare(req)

	unfollowErr := errors.New("follow error")
	service := mock.NewMockService(ctrl)
	service.EXPECT().Unfollow(ctx, gomock.Any()).Return(unfollowErr)
	controller := NewController(service)

	if assert.NoError(t, controller.Unfollow(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}
