package services

import (
	"backend/models"
	"backend/database"
)

func GetUserProfile(userID int) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, userID)
	return user, result.Error
}