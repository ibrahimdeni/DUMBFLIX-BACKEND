package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
  userRepository := repositories.RepositoryUser(mysql.DB)
  h := handlers.HandlerAuth(userRepository)

  r.HandleFunc("/register", h.Register).Methods("POST")
  r.HandleFunc("/login", h.Login).Methods("POST") // add this code
  r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}