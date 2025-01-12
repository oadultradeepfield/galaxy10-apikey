package apikey

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/model"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/service"
	"gorm.io/gorm"
)

func (ctrl *APIKeyController) GetOrCreateAPIKey(c *gin.Context) {
	currentUser, err := service.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var apiKey model.APIKey
	if err := ctrl.db.Where("id = ?", currentUser.ID).First(&apiKey).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uniqueKey := service.GenerateRandomAlphaString(64)

			apiKey = model.APIKey{
				ID:        service.GenerateRandomAlphaString(64),
				APIKey:    uniqueKey,
				UserID:    currentUser.ID,
				ExpiredAt: time.Now().Add(15 * 24 * time.Hour),
			}

			if err := ctrl.db.Create(&apiKey).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create api key"})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"data": apiKey})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch api key"})
		return
	}

	if isExpired(apiKey) {
		apiKey.APIKey = service.GenerateRandomAlphaString(64)
		apiKey.ExpiredAt = time.Now().Add(15 * 24 * time.Hour)

		if err := ctrl.db.Save(&apiKey).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update api key"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": apiKey})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": apiKey})
}

func isExpired(apiKey model.APIKey) bool {
	return apiKey.ExpiredAt.Before(time.Now())
}
