package controllers

import (
	"net/http"
	"strings"

	"holygrail-backend/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handlers struct {
	AuthDB   *gorm.DB
	Settings config.Settings
}

func New(authDB *gorm.DB, settings config.Settings) *Handlers {
	return &Handlers{AuthDB: authDB, Settings: settings}
}

// [STEP 3.1] Provide a simple public health response at the home endpoint.
func (h *Handlers) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome"})
}

func isUniqueViolation(err error) bool {
	message := strings.ToLower(err.Error())
	return strings.Contains(message, "duplicate key") || strings.Contains(message, "unique constraint")
}
