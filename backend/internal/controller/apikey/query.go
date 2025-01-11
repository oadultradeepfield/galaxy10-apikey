package apikey

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/model"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/service"
)

func (ctrl *APIKeyController) GetAllAPIKeys(c *gin.Context) {
	currentUser, err := service.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var apiKeys []model.APIKey
	if err := ctrl.db.Where("user_id = ?", currentUser.ID).Find(&apiKeys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch api keys"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": apiKeys})
}
