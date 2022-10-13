package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	filmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepository)

	r.HandleFunc("/films", h.FindFilms).Methods("GET")         //get alll
	r.HandleFunc("/film/{id}", h.GetFilm).Methods("GET")       //select
	r.HandleFunc("/film", middleware.Auth(middleware.ChekAdmin(middleware.UploadFile(h.CreateFilm)))).Methods("POST")        // add
	r.HandleFunc("/film/{id}", middleware.Auth(middleware.ChekAdmin(h.UpdateFilm))).Methods("PATCH")  // edite
	r.HandleFunc("/film/{id}", middleware.Auth(middleware.ChekAdmin(h.DeleteFilm))).Methods("DELETE") // delete

}
