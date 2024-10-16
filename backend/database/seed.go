package database

import (
	"backend/models"
)

func Seed() {
	DB.Create(&models.User{Username: "seller", Password: "password", Role: "seller"})
	DB.Create(&models.User{Username: "buyer", Password: "password", Role: "buyer"})
	DB.Create(&models.Product{Name: "Sample Product", Description: "This is a sample product", Price: 99.99, ImageURL: "http://example.com/image.png", SellerID: 1})
}