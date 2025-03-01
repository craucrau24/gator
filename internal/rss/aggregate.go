package rss

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/craucrau24/gator/internal/config"
	"github.com/craucrau24/gator/internal/database"
	"github.com/google/uuid"
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

	now := time.Now()

	for _, item := range feed.Channel.Item {
		var pubDate *time.Time
		layouts := [1]string{time.RFC1123Z}
		for _, layout := range layouts {
			time, err := time.Parse(layout, item.PubDate)
			if err == nil {
				pubDate = &time
				break
			}
		}
		if pubDate == nil {
			fmt.Printf("Error parsing date `%s`", item.PubDate)
			continue
		}

		post, err := s.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: *pubDate,
			FeedID:      next.ID,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "posts_url_key") {
				fmt.Printf("Error when post put in database: %v\n", err)
			}
			continue
		}
		fmt.Println(item.PubDate)
		fmt.Println(post)
	}
}
