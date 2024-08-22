package handler

import (
	"bytes"
	"content/internal"
	"content/internal/routes/contents"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type ReaderServiceMock struct {
	mock.Mock
}

func (r *ReaderServiceMock) GetContent(ctx context.Context) ([]contents.Content, error) {
	args := r.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]contents.Content), args.Error(1)
}

func TestGetContent(t *testing.T) {
	emptyMarshaledGetContentResponse, err := json.Marshal([]contents.Content{})
	require.NoError(t, err)

	dbErrorResponse, err := json.Marshal(internal.ErrorBase{
		Message: "database error: db error",
		Code:    5002,
	})
	require.NoError(t, err)

	t.Run("GIVEN reader returning empty contents THEN success response returned", func(t *testing.T) {
		// GIVEN
		gin.SetMode(gin.TestMode)
		httpRecorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(httpRecorder)
		ctx.Request = httptest.NewRequest(http.MethodGet, "/contents", nil)
		requestCtx := context.WithValue(ctx.Request.Context(), "request", ctx)
		ctx.Request = ctx.Request.WithContext(requestCtx)

		readerService := &ReaderServiceMock{}
		readerService.On("GetContent", mock.Anything, mock.Anything).Return([]contents.Content{}, nil)

		// WHEN
		New(readerService).GetContent(ctx)
		resp := httpRecorder.Result()
		buf := new(bytes.Buffer)

		_, err := buf.ReadFrom(resp.Body)

		// THEN
		readerService.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NoError(t, err)
		assert.JSONEq(t, string(emptyMarshaledGetContentResponse), buf.String())
	})

	t.Run("GIVEN reader returning error THEN error response returned", func(t *testing.T) {
		// GIVEN
		gin.SetMode(gin.TestMode)
		httpRecorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(httpRecorder)
		ctx.Request = httptest.NewRequest(http.MethodGet, "/contents", nil)
		requestCtx := context.WithValue(ctx.Request.Context(), "request", ctx)
		ctx.Request = ctx.Request.WithContext(requestCtx)

		readerService := &ReaderServiceMock{}
		readerService.On("GetContent", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

		// WHEN
		New(readerService).GetContent(ctx)
		resp := httpRecorder.Result()
		buf := new(bytes.Buffer)

		_, err := buf.ReadFrom(resp.Body)

		// THEN
		readerService.AssertExpectations(t)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.NoError(t, err)
		assert.JSONEq(t, string(dbErrorResponse), buf.String())
	})
}
