package main

import (
	"github.com/Antonvasilache/gator/internal/config"
	"github.com/Antonvasilache/gator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}