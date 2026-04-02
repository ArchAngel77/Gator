
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"github.com/google/uuid"
	"gator/internal/database"
	"strings"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}
	log.Printf("Collecting feeds every %s...", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feed to fetch", err)
		return
	}
	s.db.MarkFeedFetched(context.Background(), feed.ID)
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
                log.Println("Couldn't get next feed to fetch", err)
                return
        }
	for _, item := range feedData.Channel.Item {
	parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
	if err != nil {
		parsedTime, err = time.Parse(time.RFC1123, item.PubDate)
	}
	if err != nil {
		log.Printf("couldn't parse time %q: %v", item.PubDate, err)
		continue
	}
	_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
		ID:		uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Title:		sql.NullString{String: item.Title, Valid: true},
		Url:		sql.NullString{String: item.Link, Valid: true},
		Description:	sql.NullString{String: item.Description, Valid: true},
		PublishedAt:	sql.NullTime{Time: parsedTime, Valid: true},
		FeedID:		feed.ID,
	})
	if err != nil {
		if !strings.Contains(err.Error(), "duplicate key") {
			log.Printf("couldn't create post: %v", err)
			}
		}
	}
}
