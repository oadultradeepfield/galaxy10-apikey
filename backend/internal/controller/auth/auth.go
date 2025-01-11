package auth

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type AuthController struct {
	db          *gorm.DB
	oauthConfig *oauth2.Config
	frontendURL string
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		db: db,
		oauthConfig: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.profile",
				"https://www.googleapis.com/auth/userinfo.email",
			},
			Endpoint: google.Endpoint,
		},
		frontendURL: os.Getenv("FRONTEND_REDIRECT_URL"),
	}
}
