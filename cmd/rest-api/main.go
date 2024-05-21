package main

import (
	"api/internal/api"
)

func main() {
	api := api.New()
	defer api.Close()

	api.SeedWithFakeData()
	api.Run()
}
