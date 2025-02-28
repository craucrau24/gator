package rss

import (
	"context"
	"fmt"
	"time"

	"github.com/craucrau24/gator/internal/config"
	"github.com/craucrau24/gator/internal/database"
)

func ScrapeFeeds(s *config.State) {
	next, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return
	}
	s.DB.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        next.ID,
	})
	fmt.Printf("Fetching feed '%s' with url %s\n", next.Name, next.Url)
	feed, err := FetchFeed(context.Background(), next.Url)
	if err != nil {
		return
	}
	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
	}
}
