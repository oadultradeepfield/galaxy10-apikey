package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/model"
)

func GetCurrentUser(c *gin.Context) (*model.User, error) {
	currentUser, exists := c.Get("user")
	if !exists {
		return nil, errors.New("user not authenticated")
	}

	currentUserData, ok := currentUser.(*model.User)
	if !ok {
		return nil, errors.New("invalid user data")
	}

	return currentUserData, nil
}
