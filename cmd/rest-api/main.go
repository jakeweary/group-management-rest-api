package main

import (
	"api/internal/api"
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	api, err := api.New()
	if err != nil {
		return fmt.Errorf("api.New(): %v", err)
	}
	defer api.Close()

	if err := api.SeedWithFakeData(); err != nil {
		return fmt.Errorf("api.SeedWithFakeData(): %v", err)
	}

	if err := api.Run(); err != nil {
		return fmt.Errorf("api.Run(): %v", err)
	}

	return nil
}
