package utilis

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string) (string, error) {
	key := os.Getenv("JWT_SECRET")
	if len(key) < 32 {
		return "", errors.New("JWT_SECRET must be at least 32 characters")
	}

	claims := jwt.RegisteredClaims{
		Subject:   email,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	// Create a JWT using our claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	key := os.Getenv("JWT_SECRET")
	if len(key) < 32 {
		return nil, errors.New("JWT_SECRET must be at least 32 characters")
	}

	return jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithExpirationRequired(),
	)
}

// LogoutSession overwrites the access_token cookie with an expired date to log the user out.
func LogoutSession(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Expired timestamp instructs browser to delete cookie
		MaxAge:   -1,
		HTTPOnly: true,
		SameSite: "Lax",
	})
}

func SessionEmail(c *fiber.Ctx) (string, bool) {
	tokenString := c.Cookies("access_token")
	if tokenString == "" {
		authorization := strings.Fields(c.Get("Authorization"))
		if len(authorization) == 2 && strings.EqualFold(authorization[0], "Bearer") {
			tokenString = authorization[1]
		}
	}
	if tokenString == "" {
		return "", false
	}

	token, err := ValidateToken(tokenString)
	if err != nil || !token.Valid {
		return "", false
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || claims.Subject == "" {
		return "", false
	}
	return claims.Subject, true
}

// Session checks whether the request has a valid access token.
func Session(c *fiber.Ctx) bool {
	_, valid := SessionEmail(c)
	return valid
}
