package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	allowedOriginsString := os.Getenv("ALLOWED_ORIGINS")
	if allowedOriginsString == "" {
		log.Println("Warning: Allowed origins are not set in the environment variables, using default '*'")
		allowedOriginsString = "*"
	}

	allowedOrigins := strings.Split(allowedOriginsString, ",")

	return cors.New(cors.Config{
		AllowOrigins:  allowedOrigins,
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	})
}
