package main

import "github.com/Antonvasilache/gator/internal/config"

type state struct {
	cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}