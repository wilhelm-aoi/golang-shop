package services

import (
	"errors"
	"backend/models"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Authenticate(loginRequest LoginRequest) (string, error) {
	var user models.User
	result := database.DB.Where("username = ? AND password = ?", loginRequest.Username, loginRequest.Password).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("invalid credentials")
	}

	// Generate token (omitted for simplicity)
	return "mock-token", nil
}