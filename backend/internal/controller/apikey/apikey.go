package apikey

import "gorm.io/gorm"

type APIKeyController struct {
	db *gorm.DB
}

func NewAPIKeyController(db *gorm.DB) *APIKeyController {
	return &APIKeyController{db: db}
}
