package routes

import (
	. "github.com/cronjob-service/pkg/middlewares"
	"github.com/gorilla/mux"

	. "github.com/cronjob-service/pkg/controllers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// Home
	router.HandleFunc("/", MiddlewareJSON(GetHome)).Methods("GET")

	// Offer
	router.HandleFunc("/api/jobs", MiddlewareJSON(GetJobs)).Methods("GET")
	router.HandleFunc("/api/offer", MiddlewareJSON(StartOfferJob)).Methods("POST")
	router.HandleFunc("/api/offer/{id}", MiddlewareJSON(DeleteOfferJob)).Methods("DELETE")

	return router
}
