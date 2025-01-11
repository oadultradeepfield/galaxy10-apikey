package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/model"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/service"
)

func (ctrl *UserController) GetCurrentUserInfo(c *gin.Context) {
	currentUser, err := service.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := ctrl.db.Where("user_id = ?", currentUser.ID).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch current user information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
