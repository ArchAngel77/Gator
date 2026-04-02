

package main

import (
	"context"
	"fmt"
	"strconv"
	"gator/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	limit := 2
	if len(cmd.Args) >= 1 {
		parsed, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}
		limit = parsed
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}
	for _, post := range posts {
		fmt.Printf("%s\n%s\n\n", post.Title.String, post.Url.String)
	}
	return nil
}
