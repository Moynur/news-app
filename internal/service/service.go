//go:generate mockgen -package=service -source=service.go -destination=./service_mock.go Service
package service

import (
	"github.com/moynur/news-app/internal/models"
	"github.com/moynur/news-app/internal/store"
)

type Service interface {
	GetArticles(request models.GetArticlesRequest) (models.GetArticlesResponse, error)
}

type service struct {
	store      store.Storer
	rssFeedUrl string
}

func NewService(db store.Storer) *service {
	return &service{
		store: db,
	}
}
