package service

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/model"
	"golang.org/x/exp/rand"
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

func GenerateRandomAlphaString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(uint64(time.Now().UnixNano()))
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
