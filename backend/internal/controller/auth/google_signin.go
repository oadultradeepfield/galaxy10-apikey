package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/middleware"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/model"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/service"
	"golang.org/x/oauth2"
)

func (ctrl *AuthController) GoogleSignin(c *gin.Context) {
	url := ctrl.oauthConfig.AuthCodeURL("state")
	c.Redirect(http.StatusFound, url)
}

func (ctrl *AuthController) GoogleCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorization code is missing"})
		return
	}

	token, err := ctrl.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange authorization code"})
		return
	}

	userInfo, err := ctrl.getUserInfo(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user information"})
		return
	}

	user, err := ctrl.createOrUpdateUser(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process user"})
		return
	}

	accessToken, err := ctrl.generateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	redirectURL := fmt.Sprintf("%s?token=%s", ctrl.frontendURL, accessToken)
	c.Redirect(http.StatusFound, redirectURL)
}

func (ctrl *AuthController) getUserInfo(token *oauth2.Token) (*struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}, error) {
	client := ctrl.oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info from Google: %w", err)
	}
	defer resp.Body.Close()

	var userInfo struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	return &userInfo, nil
}

func (ctrl *AuthController) createOrUpdateUser(userInfo *struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}) (*model.User, error) {
	var user model.User
	if err := ctrl.db.Where("email = ?", userInfo.Email).First(&user).Error; err != nil {
		user = model.User{
			ID:        service.GenerateRandomAlphaString(64),
			Username:  userInfo.Name,
			Email:     userInfo.Email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := ctrl.db.Create(&user).Error; err != nil {
			return nil, fmt.Errorf("failed to create new user: %w", err)
		}
	}
	return &user, nil
}

func (ctrl *AuthController) generateToken(userID string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT secret is not set")
	}

	claims := &middleware.CustomClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
