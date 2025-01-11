package apikey

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/model"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/service"
	"gorm.io/gorm"
)

func (ctrl *APIKeyController) CreateAPIKey(c *gin.Context) {
	currentUser, err := service.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var input struct {
		KeyName   string     `json:"key_name" binding:"required"`
		ExpiredAt *time.Time `json:"expired_at"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var uniqueKey string
	for {
		uniqueKey = uuid.New().String()

		var existingKey model.APIKey
		err := ctrl.db.Where("api_key = ?", uniqueKey).First(&existingKey).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				break
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate unique api key"})
			return
		}
	}

	apiKey := model.APIKey{
		ID:        uuid.New().String(),
		KeyName:   input.KeyName,
		APIKey:    uniqueKey,
		UserID:    currentUser.ID,
		ExpiredAt: time.Now().AddDate(0, 0, 15),
	}

	if input.ExpiredAt != nil {
		apiKey.ExpiredAt = *input.ExpiredAt
	}

	if err := ctrl.db.Create(&apiKey).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create api key"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": apiKey})
}

func (ctrl *APIKeyController) UpdateAPIKey(c *gin.Context) {
	currentUser, err := service.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	keyID := c.Param("id")

	var input struct {
		KeyName   string     `json:"key_name"`
		ExpiredAt *time.Time `json:"expired_at"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var apiKey model.APIKey
	if err := ctrl.db.Where("id = ? AND user_id = ?", keyID, currentUser.ID).First(&apiKey).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "api key not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch api key"})
		return
	}

	updates := make(map[string]interface{})
	if input.KeyName != "" {
		updates["key_name"] = input.KeyName
	}
	if input.ExpiredAt != nil {
		updates["expired_at"] = input.ExpiredAt
	}

	if err := ctrl.db.Model(&apiKey).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update api key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": apiKey})
}

func (ctrl *APIKeyController) DeleteAPIKey(c *gin.Context) {
	currentUser, err := service.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	keyID := c.Param("id")

	result := ctrl.db.Where("id = ? AND user_id = ?", keyID, currentUser.ID).Delete(&model.APIKey{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete api key"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "api key not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "api key deleted successfully"})
}
