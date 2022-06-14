package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/moynur/news-app/internal/models"
	"github.com/moynur/news-app/internal/service"
	handler "github.com/moynur/news-app/internal/transport/http"
)

const (
	loadURL = "/loadArticles"
)

func TestHandler_NewHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ms := service.NewMockService(ctrl)
	h, err := handler.NewHandler(ms)
	assert.NoError(t, err)
	assert.NotNil(t, h)

}

func TestHandler_LoadArticles(t *testing.T) {
	t.Run("should return some articles", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ms := service.NewMockService(ctrl)

		h, err := handler.NewHandler(ms)
		assert.NoError(t, err)

		assert.NotNil(t, h)

		request := handler.LoadArticlesReq{
			Cursor:   0,
			Category: "some category",
			Provider: "some provider",
			Title:    "some title",
		}

		reqMarshalled, err := json.Marshal(request)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, loadURL, bytes.NewReader(reqMarshalled))

		expectedServerReq := models.GetArticlesRequest{
			Cursor:   request.Cursor,
			Category: request.Category,
			Provider: request.Provider,
			Title:    request.Title,
		}

		expectedServerResp := models.GetArticlesResponse{
			NextCursor: 0,
			Articles: []models.Article{
				{
					ID:       0,
					Title:    "some title",
					Summary:  "some summary",
					ImageRef: "some image url",
					Link:     "some page url",
				},
				{
					ID:       1,
					Title:    "some title",
					Summary:  "some summary",
					ImageRef: "some image url",
					Link:     "some page url",
				},
			},
		}

		ms.EXPECT().GetArticles(expectedServerReq).Return(expectedServerResp, nil)

		h.LoadArticles(w, r)
		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		err = resp.Body.Close()
		assert.NoError(t, err)

		expected := handler.LoadArticlesResp{
			NextCursor: 0, // the handler doesn't handle the cursor value just passes back what the service provides
			Articles: []handler.Article{
				{
					Title:    "some title",
					Summary:  "some summary",
					ImageRef: "some image url",
					Link:     "some page url",
				},
				{
					Title:    "some title",
					Summary:  "some summary",
					ImageRef: "some image url",
					Link:     "some page url",
				},
			},
		}

		var out handler.LoadArticlesResp
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&out))
		assert.Equal(t, expected, out)
	})

	t.Run("should handle an error when not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ms := service.NewMockService(ctrl)

		h, err := handler.NewHandler(ms)
		assert.NoError(t, err)

		assert.NotNil(t, h)

		request := handler.LoadArticlesReq{
			Cursor:   0,
			Category: "some category",
			Provider: "some provider",
			Title:    "some title",
		}

		reqMarshalled, err := json.Marshal(request)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, loadURL, bytes.NewReader(reqMarshalled))

		expectedServerReq := models.GetArticlesRequest{
			Cursor:   request.Cursor,
			Category: request.Category,
			Provider: request.Provider,
			Title:    request.Title,
		}

		expectedServerResp := models.GetArticlesResponse{}

		ms.EXPECT().GetArticles(expectedServerReq).Return(expectedServerResp, service.ErrNotFound)

		h.LoadArticles(w, r)
		resp := w.Result()
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("should handle a generic error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ms := service.NewMockService(ctrl)

		h, err := handler.NewHandler(ms)
		assert.NoError(t, err)

		assert.NotNil(t, h)

		request := handler.LoadArticlesReq{
			Cursor:   0,
			Category: "some category",
			Provider: "some provider",
			Title:    "some title",
		}

		reqMarshalled, err := json.Marshal(request)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, loadURL, bytes.NewReader(reqMarshalled))

		expectedServerReq := models.GetArticlesRequest{
			Cursor:   request.Cursor,
			Category: request.Category,
			Provider: request.Provider,
			Title:    request.Title,
		}

		expectedServerResp := models.GetArticlesResponse{}

		ms.EXPECT().GetArticles(expectedServerReq).Return(expectedServerResp, errors.New("something failed"))

		h.LoadArticles(w, r)
		resp := w.Result()
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("should return error when request is bad", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ms := service.NewMockService(ctrl)

		h, err := handler.NewHandler(ms)
		assert.NoError(t, err)

		assert.NotNil(t, h)

		request := handler.LoadArticlesResp{
			NextCursor: 0,
			Articles: []handler.Article{
				{
					Title:    "",
					Summary:  "",
					ImageRef: "",
					Link:     "",
				},
			},
		}

		reqMarshalled, err := json.Marshal(request)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, loadURL, bytes.NewReader(reqMarshalled))

		ms.EXPECT().GetArticles(gomock.Any()).Times(0)

		h.LoadArticles(w, r)
		resp := w.Result()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		err = resp.Body.Close()
		assert.NoError(t, err)
	})
}
