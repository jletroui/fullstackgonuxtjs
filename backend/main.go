package main

import (
	"backend/config"
	"log"
)

func main() {
	config, err := config.LoadFromEnv("dev")
	if err != nil {
		log.Fatalf("Cannot load config: %s", err)
	}
	log.Printf("Database: %s\n", config.PostgresDatabase)
}
