package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {

	UserRoutes(r)
	FilmRoutes(r)
	EpisodeRoutes(r)
	AuthRoutes(r)
	CategoryRoutes(r)
	TransactionRoutes(r)
}
