package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Antonvasilache/gator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("reset command expects two arguments. Usage: %s <name> <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		Name: name,
		Url: url,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed follow: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("Feed followed successfully:")
	PrintFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("===================================")

	return nil
}

func printFeed(feed database.Feed){
	fmt.Printf(" * ID: 			%s\n", feed.ID)
	fmt.Printf(" * Created: 	%s\n", feed.CreatedAt)
	fmt.Printf(" * Updated: 	%s\n", feed.UpdatedAt)
	fmt.Printf(" * Name: 		%s\n", feed.Name)
	fmt.Printf(" * URL: 		%s\n", feed.Url)
	fmt.Printf(" * User ID: 	%s\n", feed.UserID)
}

func handleGetFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("feeds command expects no arguments. Usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not retrieve feeds: %w", err)
	}

	for _, feed := range feeds {
		username, err := s.db.GetUserNameById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("could not retrieve user: %w", err)
		}

		fmt.Printf("* Name: 		%s\n", feed.Name)
		fmt.Printf("* URL: 		%s\n", feed.Url)
		fmt.Printf("* User name: 	%s\n", username)
	}

	return nil
}