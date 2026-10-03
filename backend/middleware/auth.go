package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
	"time"

	"holygrail-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Claims struct {
	jwt.RegisteredClaims
}

func IssueJWT(email, secret string, expiration time.Duration) (string, time.Time, error) {
	if len(secret) < 32 {
		return "", time.Time{}, errors.New("JWT_SECRET must be at least 32 characters")
	}
	if expiration <= 0 {
		return "", time.Time{}, errors.New("JWT expiration must be positive")
	}
	expiresAt := time.Now().Add(expiration)
	claims := Claims{RegisteredClaims: jwt.RegisteredClaims{
		Subject: email, IssuedAt: jwt.NewNumericDate(time.Now()), ExpiresAt: jwt.NewNumericDate(expiresAt),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encoded, err := token.SignedString([]byte(secret))
	return encoded, expiresAt, err
}

func TokenHash(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func Auth(authDB *gorm.DB, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := strings.Fields(c.GetHeader("Authorization"))
		if len(authorization) != 2 || !strings.EqualFold(authorization[0], "Bearer") {
			unauthorized(c)
			return
		}
		rawToken := authorization[1]
		// [STEP 3.4] Verify the JWT signature, HS256 algorithm, expiration, and account subject before trusting the session.
		claims := &Claims{}
		parsed, err := jwt.ParseWithClaims(rawToken, claims, func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}), jwt.WithExpirationRequired())
		if err != nil || !parsed.Valid || claims.Subject == "" {
			unauthorized(c)
			return
		}

		var session models.SessionToken
		err = authDB.Where("token_hash = ? AND user_email = ? AND revoked_at IS NULL AND expires_at > ?",
			TokenHash(rawToken), claims.Subject, time.Now()).First(&session).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			unauthorized(c)
			return
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not validate session"})
			return
		}

		// [STEP 3.4] Require an account record as well as a valid JWT and active server-side session.
		var account models.AuthCredential
		err = authDB.Where("email = ?", claims.Subject).First(&account).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var legacyAdmin models.AdminUser
			err = authDB.Where("email = ? AND is_primary = ?", claims.Subject, true).First(&legacyAdmin).Error
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			unauthorized(c)
			return
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not validate account"})
			return
		}

		c.Set("user_email", claims.Subject)
		c.Set("session_token_hash", session.TokenHash)
		c.Next()
	}
}

func unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "A valid session is required"})
}
