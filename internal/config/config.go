package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port                string
	DatabaseURL         string
	CORSAllowedOrigins  []string
	LineLIFFID          string
	LineChannelToken    string
	LineChannelSecret   string
	AdminJWTSecret      string
	AdminJWTExpiryHours int
	AdminSeedUsername   string
	AdminSeedPassword   string
	LineAuthBypass      bool
}

func Load() Config {
	return Config{
		Port:                getEnv("PORT", "8080"),
		DatabaseURL:         getEnv("DATABASE_URL", ""),
		CORSAllowedOrigins:  splitCSV(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173,http://localhost:5174")),
		LineLIFFID:          getEnv("LINE_LIFF_ID", ""),
		LineChannelToken:    getEnv("LINE_CHANNEL_ACCESS_TOKEN", ""),
		LineChannelSecret:   getEnv("LINE_CHANNEL_SECRET", ""),
		AdminJWTSecret:      getEnv("ADMIN_JWT_SECRET", ""),
		AdminJWTExpiryHours: getEnvInt("ADMIN_JWT_EXPIRY_HOURS", 24),
		AdminSeedUsername:   getEnv("ADMIN_SEED_USERNAME", "admin"),
		AdminSeedPassword:   getEnv("ADMIN_SEED_PASSWORD", ""),
		LineAuthBypass:      getEnv("LINE_AUTH_BYPASS", "false") == "true",
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}

func splitCSV(v string) []string {
	parts := strings.Split(v, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
