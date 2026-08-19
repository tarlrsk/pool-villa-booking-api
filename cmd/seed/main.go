package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/deday-pool-villa/backend/internal/config"
	"github.com/deday-pool-villa/backend/internal/db"
)

// Runs seeding only — the default day-rate grid and the seed admin account,
// each only inserted if their table is empty. Assumes migrations have
// already been run (run cmd/migrate first on a fresh database).
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on process environment")
	}

	cfg := config.Load()
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	gdb := db.Connect(cfg)
	db.Seed(gdb, cfg)
	log.Println("seeding complete")
}
