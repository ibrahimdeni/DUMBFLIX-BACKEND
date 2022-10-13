package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func EpisodeRoutes(r *mux.Router) {
	EpisodeRepository := repositories.RepositoryEpisode(mysql.DB)
	h := handlers.HandlerEpisode(EpisodeRepository)

	r.HandleFunc("/episodes", h.FindEpisodes).Methods("GET")         //get alll
	r.HandleFunc("/episode/{id}", h.GetEpisode).Methods("GET")       //select
	r.HandleFunc("/episode", h.CreateEpisode).Methods("POST")        // add
	r.HandleFunc("/episode/{id}", h.UpdateEpisode).Methods("PATCH")  // edite
	r.HandleFunc("/episode/{id}", h.DeleteEpisode).Methods("DELETE") // delete

}
