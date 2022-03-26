package post

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
	"strider-backend-test.com/api/routes/post/mock"
	"strings"
	"testing"
)

func TestCtlCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	b, _ := json.Marshal(&mock.PostRequest)
	req := httptest.NewRequest(http.MethodPost, "/v1/strider/posts", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, ctx := prepare(req)

	service := mock.NewMockService(ctrl)
	service.EXPECT().Create(ctx, gomock.Any()).Return(&mock.Post, nil)
	controller := NewController(service)

	if assert.NoError(t, controller.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestCtlCreateBindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest(http.MethodPost, "/v1/strider/posts", strings.NewReader(mock.BrokenRequest))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, _ := prepare(req)

	service := mock.NewMockService(ctrl)
	controller := NewController(service)

	if assert.NoError(t, controller.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCtlCreateValidateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	b, _ := json.Marshal(mock.PostRequestWithoutCreator)
	req := httptest.NewRequest(http.MethodPost, "/v1/strider/posts", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, _ := prepare(req)

	service := mock.NewMockService(ctrl)
	controller := NewController(service)

	if assert.NoError(t, controller.Create(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCtlCreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	b, _ := json.Marshal(&mock.Post)
	req := httptest.NewRequest(http.MethodPost, "/v1/strider/posts", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec, c, ctx := prepare(req)

	service := mock.NewMockService(ctrl)
	service.EXPECT().Create(ctx, gomock.Any()).Return(nil, errors.New("create error"))
	controller := NewController(service)

	if assert.NoError(t, controller.Create(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}

func TestCtlList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest(http.MethodGet, "/v1/strider/posts?user=de1753e0-bf73-4813-9e0b-5e7cea8fa008", nil)

	rec, c, ctx := prepare(req)

	service := mock.NewMockService(ctrl)
	service.EXPECT().List(ctx, gomock.Any()).Return(mock.Posts, nil)
	controller := NewController(service)

	if assert.NoError(t, controller.List(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestCtlListBindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest(http.MethodGet, "/v1/strider/posts?offset=offset", nil)

	e := echo.New()
	e.Binder = middleware.NewBinder()
	e.Validator = middleware.NewValidator()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service := mock.NewMockService(ctrl)
	controller := NewController(service)

	if assert.NoError(t, controller.List(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCtlListValidateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest(http.MethodGet, "/v1/strider/posts?user=1234&follower=1234", nil)

	e := echo.New()
	e.Binder = middleware.NewBinder()
	e.Validator = middleware.NewValidator()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service := mock.NewMockService(ctrl)
	controller := NewController(service)

	if assert.NoError(t, controller.List(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCtlListError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest(http.MethodGet, "/v1/strider/posts?user=de1753e0-bf73-4813-9e0b-5e7cea8fa008", nil)

	rec, c, ctx := prepare(req)

	service := mock.NewMockService(ctrl)
	service.EXPECT().List(ctx, gomock.Any()).Return(nil, errors.New("list error"))
	controller := NewController(service)

	if assert.NoError(t, controller.List(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}

func TestCtlEmptyList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := httptest.NewRequest(http.MethodGet, "/v1/strider/posts?user=de1753e0-bf73-4813-9e0b-5e7cea8fa008", nil)

	rec, c, ctx := prepare(req)

	service := mock.NewMockService(ctrl)
	service.EXPECT().List(ctx, gomock.Any()).Return(mock.EmptyPosts, nil)
	controller := NewController(service)

	if assert.NoError(t, controller.List(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
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
