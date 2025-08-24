package main

import (
	"fmt"
	"log"

	"github.com/theHousedev/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	fmt.Println(cfg)

	cfg, err = cfg.SetUser("house")
	if err != nil {
		log.Fatalf("Failed to set user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	fmt.Println(cfg)
}
