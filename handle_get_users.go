package main

import (
	"context"
	"fmt"
)

func handleGetUsers (s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("users command expects no arguments. Usage: %s", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not retrieve users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}