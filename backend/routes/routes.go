package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"backend/controllers"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/user/profile", controllers.GetUserProfile).Methods("GET")

	return r
}