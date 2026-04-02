
package main

import(
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/google/uuid"
	"gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("a url is required")
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams {
		ID:		uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		UserID:		user.ID,
		FeedID:		feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(follow.FeedName)
	fmt.Println(follow.UserName)
	return nil
}
func handlerFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, ff := range follows {
	fmt.Println(ff.FeedName)
	}
	return nil
}
func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
                return errors.New("a url is required")
        }
        url := cmd.Args[0]
        feed, err := s.db.GetFeedByURL(context.Background(), url)
        if err != nil {
                return err
        }
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams {
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("%s unfollowed successfully!\n", feed.Name)
	return nil
}
