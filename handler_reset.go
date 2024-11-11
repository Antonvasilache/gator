package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("reset command expects no arguments. Usage: %s", cmd.Name)
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not reset users: %w", err)
	}

	fmt.Println("users successfully reset")

	return nil;
}