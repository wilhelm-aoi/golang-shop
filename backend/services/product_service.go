package services

import (
	"backend/models"
	"backend/database"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := database.DB.Find(&products)
	return products, result.Error
}