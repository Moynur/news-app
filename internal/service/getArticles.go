package service

import (
	"errors"
	"log"

	"github.com/moynur/news-app/internal/models"
	"github.com/moynur/news-app/internal/store"
)

var (
	ErrNotFound = errors.New("no articles found matching criteria")
)

func (s *service) GetArticles(req models.GetArticlesRequest) (models.GetArticlesResponse, error) {
	log.Println("reached service")
	var response models.GetArticlesResponse
	// number of records could be a config or a parameter from the client request
	articles, err := s.store.GetRecordsAfterID(req.Cursor, 3, store.Filters{
		// haven't implemented others but this is to showcase how the filters work
		Title:         req.Title,
		Description:   "",
		Link:          "",
		Category:      req.Category,
		CreatedAfter:  nil,
		CreatedBefore: nil,
	})
	if err != nil {
		return response, err
	}
	if len(articles) == 0 {
		return response, ErrNotFound
	}
	for _, article := range articles {
		response.Articles = append(response.Articles, models.Article{
			ID:       int(article.ID),
			Title:    article.Title,
			Summary:  article.Description,
			ImageRef: article.Thumbnail,
			Link:     article.Link,
		})
	}
	response.NextCursor = response.Articles[len(response.Articles)-1].ID
	log.Println("response", response.NextCursor)
	return response, nil
}
