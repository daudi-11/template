package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TestCORSAllowsConfiguredOriginAndPreflight(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(CORS())
	router.GET("/", func(c *gin.Context) { c.Status(http.StatusOK) })

	request := httptest.NewRequest(http.MethodOptions, "/", nil)
	request.Header.Set("Origin", "http://localhost:5173")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if response.Code != http.StatusNoContent {
		t.Fatalf("preflight status = %d, want %d", response.Code, http.StatusNoContent)
	}
	if got := response.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:5173" {
		t.Fatalf("allow origin = %q", got)
	}
	if got := response.Header().Get("Access-Control-Allow-Credentials"); got != "true" {
		t.Fatalf("allow credentials = %q", got)
	}
}

func TestCORSDoesNotReflectUnknownOrigin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(CORS())
	router.GET("/", func(c *gin.Context) { c.Status(http.StatusOK) })

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Origin", "https://untrusted.example")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if got := response.Header().Get("Access-Control-Allow-Origin"); got != "" {
		t.Fatalf("unexpected allow origin: %q", got)
	}
}

func TestCORSAllowsReactAndLocalhostOrigins(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(CORS())
	router.GET("/", func(c *gin.Context) { c.Status(http.StatusOK) })

	for _, origin := range []string{"http://localhost:5173", "http://localhost:3000"} {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set("Origin", origin)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		if got := response.Header().Get("Access-Control-Allow-Origin"); got != origin {
			t.Errorf("origin %q received allow-origin %q", origin, got)
		}
	}
}

func TestIssueJWTRequiresStrongSecretAndSetsExpiration(t *testing.T) {
	if _, _, err := IssueJWT("user@example.com", "short", time.Hour); err == nil {
		t.Fatal("IssueJWT() expected an error for a short secret")
	}

	secret := "test-secret-with-at-least-32-characters-long"
	token, expiresAt, err := IssueJWT("user@example.com", secret, time.Hour)
	if err != nil {
		t.Fatalf("IssueJWT() error = %v", err)
	}
	if token == "" || time.Until(expiresAt) <= 0 || time.Until(expiresAt) > time.Hour {
		t.Fatalf("unexpected token or expiration: token-set=%t expiration=%v", token != "", expiresAt)
	}
	claims := &Claims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (any, error) { return []byte(secret), nil })
	if err != nil || !parsed.Valid || claims.Subject != "user@example.com" {
		t.Fatalf("JWT subject = %q, parse error = %v", claims.Subject, err)
	}
}
