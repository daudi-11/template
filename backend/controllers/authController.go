package controllers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"holygrail-backend/middleware"
	"holygrail-backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type credentialsRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// [STEP 3.1] HTTP handler for Register.
func (h *Handlers) Register(c *gin.Context) {
	var request credentialsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Send valid JSON with name, email, and password"})
		return
	}
	request.Name = strings.TrimSpace(request.Name)
	request.Email = strings.TrimSpace(request.Email)
	if request.Name == "" || request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Name, email, and password are required"})
		return
	}

	var existingAdmin models.AdminUser
	if err := h.AuthDB.Where("email = ?", request.Email).First(&existingAdmin).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "An account with this email already exists"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not validate email"})
		return
	}

	var existingAccount models.AuthCredential
	if err := h.AuthDB.Where("email = ?", request.Email).First(&existingAccount).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "An account with this email already exists"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not validate email"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to secure password"})
		return
	}
	credential := models.AuthCredential{Name: request.Name, Email: request.Email, PasswordHash: string(hash), AccountType: "user"}
	if err := h.AuthDB.Create(&credential).Error; err != nil {
		if isUniqueViolation(err) {
			c.JSON(http.StatusConflict, gin.H{"message": "An account with this email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Account created successfully", "user": gin.H{"name": credential.Name, "email": credential.Email}})
}

// [STEP 3.1] HTTP handler for Login.
func (h *Handlers) Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Send valid JSON with email and password"})
		return
	}
	request.Email = strings.TrimSpace(request.Email)
	if request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email and password are required"})
		return
	}

	account, err := h.authenticate(request.Email, request.Password)
	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Login failed"})
		return
	}

	expiration := time.Duration(h.Settings.JWTExpirationHours) * time.Hour
	token, expiresAt, err := middleware.IssueJWT(request.Email, h.Settings.JWTSecret, expiration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create token"})
		return
	}
	session := models.SessionToken{
		TokenHash: middleware.TokenHash(token), UserEmail: request.Email, ExpiresAt: expiresAt,
	}
	if err := h.AuthDB.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "user": gin.H{"name": account.Name, "email": account.Email}})
}

func (h *Handlers) authenticate(email, password string) (models.AuthCredential, error) {
	var admin models.AdminUser
	if err := h.AuthDB.Where("email = ?", email).First(&admin).Error; err == nil {
		if bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password)) != nil {
			return models.AuthCredential{}, bcrypt.ErrMismatchedHashAndPassword
		}
		if admin.IsPrimary {
			return models.AuthCredential{Name: admin.Name, Email: admin.Email, PasswordHash: admin.PasswordHash}, nil
		}
		return models.AuthCredential{}, gorm.ErrRecordNotFound
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.AuthCredential{}, err
	}

	var credential models.AuthCredential
	if err := h.AuthDB.Where("email = ?", email).First(&credential).Error; err != nil {
		return models.AuthCredential{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(credential.PasswordHash), []byte(password)); err != nil {
		return models.AuthCredential{}, err
	}
	return credential, nil   
}

// [STEP 3.1] HTTP handler for Logout.
func (h *Handlers) Logout(c *gin.Context) {
	tokenHash := c.GetString("session_token_hash")
	if tokenHash != "" {
		now := time.Now()
		if err := h.AuthDB.Model(&models.SessionToken{}).Where("token_hash = ? AND revoked_at IS NULL", tokenHash).Update("revoked_at", now).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not end session"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
