package service_test

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/moynur/news-app/internal/models"
	"github.com/moynur/news-app/internal/service"
	"github.com/moynur/news-app/internal/store"
)

func Test_service_GetArticles(t *testing.T) {
	type args struct {
		req      models.GetArticlesRequest
		resp     []store.NewsArticle
		storeErr error
	}
	tests := []struct {
		name    string
		args    args
		want    models.GetArticlesResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "correctly maps articles",
			args: args{
				req: models.GetArticlesRequest{
					Cursor:   0,
					Category: "",
					Provider: "",
					Title:    "",
				},
				resp: []store.NewsArticle{
					{
						ID:          0,
						Title:       "someTitle",
						Description: "someDescription",
						Link:        "someLink",
						Category:    "someCategory",
						Thumbnail:   "someThumbnail",
						CreatedAt:   time.Now(),
					},
					{
						ID:          1,
						Title:       "someTitle",
						Description: "someDescription",
						Link:        "someLink",
						Category:    "someCategory",
						Thumbnail:   "someThumbnail2",
						CreatedAt:   time.Now(),
					},
				},
				storeErr: nil,
			},
			want: models.GetArticlesResponse{
				NextCursor: 1,
				Articles: []models.Article{
					{
						ID:       0,
						Title:    "someTitle",
						Summary:  "someDescription",
						ImageRef: "someThumbnail",
						Link:     "someLink",
					},
					{
						ID:       1,
						Title:    "someTitle",
						Summary:  "someDescription",
						ImageRef: "someThumbnail2",
						Link:     "someLink",
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "returns error when no articles found",
			args: args{
				req: models.GetArticlesRequest{
					Cursor:   0,
					Category: "",
					Provider: "",
					Title:    "",
				},
				resp:     []store.NewsArticle{},
				storeErr: nil,
			},
			want:    models.GetArticlesResponse{},
			wantErr: assert.Error,
		},
		{
			name: "returns error when store errors",
			args: args{
				req: models.GetArticlesRequest{
					Cursor:   0,
					Category: "",
					Provider: "",
					Title:    "",
				},
				resp:     []store.NewsArticle{},
				storeErr: errors.New("cant connect to db"),
			},
			want:    models.GetArticlesResponse{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			ms := store.NewMockStorer(ctrl)
			s := service.NewService(ms)
			ms.EXPECT().GetRecordsAfterID(tt.args.req.Cursor, 3, store.Filters{
				Title:         tt.args.req.Title,
				Description:   "",
				Link:          "",
				Category:      tt.args.req.Category,
				CreatedAfter:  nil,
				CreatedBefore: nil,
			}).Return(tt.args.resp, tt.args.storeErr)
			got, err := s.GetArticles(tt.args.req)
			log.Println("got smthn", got)
			if !tt.wantErr(t, err, fmt.Sprintf("GetArticles(%v)", tt.args.req)) {
				return
			}
			log.Println("check this")
			assert.Equalf(t, tt.want, got, "GetArticles(%v)", tt.args.req)
		})
	}
}
