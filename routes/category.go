package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	r.HandleFunc("/categories", h.FindCategories).Methods("GET")         //get alll
	r.HandleFunc("/category/{id}", h.GetCategory).Methods("GET")       //select
	r.HandleFunc("/category", h.CreateCategory).Methods("POST")        // add
	r.HandleFunc("/category/{id}", h.UpdateCategory).Methods("PATCH")  // edite
	r.HandleFunc("/category/{id}", h.DeleteCategory).Methods("DELETE") // delete

}