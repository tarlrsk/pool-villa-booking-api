package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/deday-pool-villa/backend/internal/config"
	"github.com/deday-pool-villa/backend/internal/db"
)

// Runs schema migration only — no seeding, no server.
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on process environment")
	}

	cfg := config.Load()
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	gdb := db.Connect(cfg)
	if err := db.Migrate(gdb); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("migration complete")
}
