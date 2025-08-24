package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commandSet map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandSet[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commandSet[cmd.name]
	if !ok {
		return fmt.Errorf("Unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func handlerLogin(s *state, cmd command) error {
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
