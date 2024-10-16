package controllers

import (
	"net/http"
	"encoding/json"
	"backend/services"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	user, err := services.GetUserProfile(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}