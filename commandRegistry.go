package main

import (
	"fmt"
)

type Command struct {
	name string
	args []string
}

type Commands struct {
	cmd map[string]func(*State, Command) error
}

func (c *Commmands)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		fmt.Println("Usage: gator login <username>")
		return fmt.Errorf("Usage: gator login <username>")
	}

	cfg, err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Login failed: %w", err)
	}

	fmt.Println("Logged in as '", cfg.Username, "'")
	return nil
}
