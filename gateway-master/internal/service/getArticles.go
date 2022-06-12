package service

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"

	"github.com/moynur/gateway/internal/models"
)

func (s *service) GetArticles(auth models.GetArticlesRequest) (models.GetArticlesResponse, error) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(s.rssFeedUrl)
	fmt.Println(feed.Title)

	log.Println("calling store")
	return models.GetArticlesResponse{
		NextCursor: 0,
		Articles:   nil,
	}, nil
}
