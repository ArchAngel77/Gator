
package main

import(
        "errors"
        "fmt"
        "context"
        "time"
        "github.com/google/uuid"
        "gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
        if len(cmd.Args) < 2 {
                return errors.New("name and url is required")
        }
	name := cmd.Args[0]
	url := cmd.Args[1]
	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:		uuid.New(),
			CreatedAt:	time.Now(),
			UpdatedAt:	time.Now(),
			Name:		name,
			Url:		url,
			UserID:		user.ID,
	},
	)
	if err != nil {
		return err
	}
        fmt.Println(feed)
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:		uuid.New(),
		CreatedAt:      time.Now(),
                UpdatedAt:      time.Now(),
                UserID:         user.ID,
		FeedID:		feed.ID,
	})
	if err != nil {
		return err
	}
        return nil
}
func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
	feedUsers, err := s.db.GetUserById(context.Background(), feed.UserID)
	if err != nil {
		return err
	}
	fmt.Printf("* Name: %s\n",feed.Name)
	fmt.Printf("* URL: %s\n",feed.Url)
	fmt.Printf("* User: %s\n",feedUsers.Name)
	}
	return nil
}
