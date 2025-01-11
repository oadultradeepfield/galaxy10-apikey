package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/database"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/route"
)

func main() {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	db := database.InitDB()
	r := gin.Default()
	route.InitRoutes(r, db)

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
