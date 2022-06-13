package service

import (
	"log"
	"time"

	"github.com/mmcdole/gofeed"

	"github.com/moynur/gateway/internal/models"
	"github.com/moynur/gateway/internal/store"
)

func (s *service) GetArticles(req models.GetArticlesRequest) (models.GetArticlesResponse, error) {
	log.Println("reached service")
	return models.GetArticlesResponse{
		NextCursor: 0,
		Articles:   nil,
	}, nil
}

func isNotValid(article *gofeed.Item) bool {
	return article.GUID == "" || article.Title == ""
}

func (s *service) RefreshArticles(rate int) {
	ticker := time.NewTicker(time.Second * time.Duration(rate)).C
	for {
		select {
		case <-ticker:
			log.Println("ticking!")
			s.LoadAndStoreArticles()
		}
	}
}

func (s *service) LoadAndStoreArticles() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(s.rssFeedUrl)
	if err != nil {
		log.Printf("unable to load articles %e", err)
	}
	for _, article := range feed.Items {
		if isNotValid(article) {
			continue
		}
		log.Println("something 2")
		if article.Image == nil || article.Image.URL == "" {
			log.Println("something 3")
			article.Image = &gofeed.Image{
				URL: "https://pbs.twimg.com/profile_images/1140654461603287040/bUUAgDF6_400x400.jpg",
			}
		}
		log.Println("something 4", article)
		err := s.store.CreateArticleIfNotExists(store.NewsArticle{
			Title:       article.Title,
			Description: article.Description,
			Link:        article.GUID,
			Thumbnail:   article.Image.URL,
			CreatedAt:   time.Now(),
		})
		if err != nil {
			log.Println("error creating an article", err)
		}
	}
}
