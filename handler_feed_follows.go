package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Antonvasilache/gator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("follow command expects one arguments. Usage: %s <feed_url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	url := cmd.Args[0]
	
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not retrieve feed: %w", err)
	}

	ffRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed follow: %w", err)
	}
	fmt.Println("Feed follow created")
	PrintFeedFollow(ffRow.UserName, ffRow.FeedName)
	return nil
}

func PrintFeedFollow(username, feedname string){
	fmt.Printf("* User: 		%s\n", username)	
	fmt.Printf("* Feed: 		%s\n", feedname)	
}

func handleListFeedFollows(s *state, cmd command) error{
	if len(cmd.Args) != 0 {
		return fmt.Errorf("following command expects no arguments. Usage: %s", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows were found")
		return nil
	}

	fmt.Printf("Feed follows for user: %s \n", user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("* %s\n", ff.FeedName)
	}

	return nil
}