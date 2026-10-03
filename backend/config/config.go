package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Settings struct {
	Port               string
	AuthDBURL          string
	JWTSecret          string
	JWTExpirationHours int
	FrontendOrigin     string
}

func Load() (Settings, error) {
	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		return Settings{}, fmt.Errorf("load .env: %w", err)
	}

	settings := Settings{
		Port:           strings.TrimSpace(os.Getenv("PORT")),
		AuthDBURL:      strings.TrimSpace(os.Getenv("AUTH_DB_URL")),
		JWTSecret:      strings.TrimSpace(os.Getenv("JWT_SECRET")),
		FrontendOrigin: strings.TrimSpace(os.Getenv("FRONTEND_ORIGIN")),
	}
	if settings.Port == "" {
		settings.Port = "3030"
	}
	port, err := strconv.Atoi(settings.Port)
	if err != nil || port < 1 || port > 65535 {
		return Settings{}, errors.New("PORT must be a valid TCP port")
	}
	if settings.AuthDBURL == "" {
		return Settings{}, errors.New("AUTH_DB_URL is required")
	}
	if len(settings.JWTSecret) < 32 {
		return Settings{}, errors.New("JWT_SECRET must be at least 32 characters")
	}

	expirationHours := strings.TrimSpace(os.Getenv("JWT_EXPIRATION_HOURS"))
	if expirationHours == "" {
		expirationHours = "1"
	}
	settings.JWTExpirationHours, err = strconv.Atoi(expirationHours)
	if err != nil || settings.JWTExpirationHours < 1 {
		return Settings{}, errors.New("JWT_EXPIRATION_HOURS must be a positive integer")
	}
	if settings.FrontendOrigin == "" {
		settings.FrontendOrigin = "http://localhost:5173"
	}

	return settings, nil
}
