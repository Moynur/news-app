package feed

import (
	"log"
	"time"

	"github.com/mmcdole/gofeed"

	"github.com/moynur/news-app/internal/store"
)

// This could all potentially be another service with a separate database
type feeder struct {
	store      store.Storer
	rssFeedURL string
}

func NewFeeder(db store.Storer, rssFeedUrl string) *feeder {
	return &feeder{
		store:      db,
		rssFeedURL: rssFeedUrl,
	}
}

func isNotValid(article *gofeed.Item) bool {
	return article.GUID == "" || article.Title == ""
}

func (s *feeder) RefreshArticles(rate int) {
	ticker := time.NewTicker(time.Second * time.Duration(rate)).C
	for {
		select {
		case <-ticker:
			log.Println("ticking!")
			s.LoadAndStoreArticles()
		}
	}
}

func (s *feeder) LoadAndStoreArticles() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(s.rssFeedURL)
	if err != nil {
		log.Printf("unable to load articles %e", err)
	}
	for _, article := range feed.Items {
		if isNotValid(article) {
			continue
		}
		if article.Image == nil || article.Image.URL == "" {
			article.Image = &gofeed.Image{
				URL: "https://pbs.twimg.com/profile_images/1140654461603287040/bUUAgDF6_400x400.jpg",
			}
		}
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
