package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	setRequiredEnvironment(t)
	t.Setenv("PORT", "")
	t.Setenv("JWT_EXPIRATION_HOURS", "")
	t.Setenv("FRONTEND_ORIGIN", "")

	settings, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if settings.Port != "3030" || settings.JWTExpirationHours != 1 {
		t.Fatalf("unexpected defaults: port=%q expiration=%d", settings.Port, settings.JWTExpirationHours)
	}
	if settings.FrontendOrigin != "http://localhost:5173" {
		t.Fatalf("unexpected frontend origin: %q", settings.FrontendOrigin)
	}
}

func TestLoadRequiresAuthDatabaseURL(t *testing.T) {
	setRequiredEnvironment(t)
	t.Setenv("AUTH_DB_URL", "")

	if _, err := Load(); err == nil {
		t.Fatal("Load() expected an error when AUTH_DB_URL is missing")
	}
}

func TestLoadRejectsInvalidExpiration(t *testing.T) {
	setRequiredEnvironment(t)
	t.Setenv("JWT_EXPIRATION_HOURS", "0")

	if _, err := Load(); err == nil {
		t.Fatal("Load() expected an error for non-positive JWT expiration")
	}
}

func setRequiredEnvironment(t *testing.T) {
	t.Helper()
	t.Setenv("AUTH_DB_URL", "postgres://auth:secret@localhost/auth")
	t.Setenv("JWT_SECRET", "test-secret-with-at-least-32-characters-long")
	t.Setenv("PORT", "3030")
	t.Setenv("JWT_EXPIRATION_HOURS", "24")
	t.Setenv("FRONTEND_ORIGIN", "http://localhost:5173")
}
