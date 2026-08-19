package db

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/deday-pool-villa/backend/internal/config"
	"github.com/deday-pool-villa/backend/internal/domain"
)

var defaultDayRates = []domain.DayRate{
	{Day: "Monday", Price: 3000},
	{Day: "Tuesday", Price: 3000},
	{Day: "Wednesday", Price: 3000},
	{Day: "Thursday", Price: 3000},
	{Day: "Friday", Price: 4500},
	{Day: "Saturday", Price: 5000},
	{Day: "Sunday", Price: 4000},
}

// Connect opens the database connection only — no migration, no seeding.
// Callers that need those run Migrate and/or Seed explicitly.
func Connect(cfg config.Config) *gorm.DB {
	gdb, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return gdb
}

// Migrate runs AutoMigrate for every domain entity. Safe to run repeatedly —
// it only adds tables/columns, never drops or renames.
func Migrate(gdb *gorm.DB) error {
	return gdb.AutoMigrate(
		&domain.Booking{},
		&domain.BlockedDate{},
		&domain.CustomPeriod{},
		&domain.DayRate{},
		&domain.Admin{},
	)
}

// Seed inserts the default day-rate grid and the seed admin account, but
// only when those tables are empty — safe to run repeatedly.
func Seed(gdb *gorm.DB, cfg config.Config) {
	var dayRateCount int64
	gdb.Model(&domain.DayRate{}).Count(&dayRateCount)
	if dayRateCount == 0 {
		if err := gdb.Create(&defaultDayRates).Error; err != nil {
			log.Printf("warning: failed to seed day rates: %v", err)
		} else {
			log.Println("seeded default day rates")
		}
	}

	var adminCount int64
	gdb.Model(&domain.Admin{}).Count(&adminCount)
	if adminCount == 0 {
		if cfg.AdminSeedPassword == "" {
			log.Println("warning: no admin account exists and ADMIN_SEED_PASSWORD is not set — admin login will be unavailable until an admin row is created")
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(cfg.AdminSeedPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("warning: failed to hash seed admin password: %v", err)
			return
		}
		admin := domain.Admin{Username: cfg.AdminSeedUsername, PasswordHash: string(hash)}
		if err := gdb.Create(&admin).Error; err != nil {
			log.Printf("warning: failed to seed admin account: %v", err)
			return
		}
		log.Printf("seeded admin account %q — change the password after first login", cfg.AdminSeedUsername)
	}
}
