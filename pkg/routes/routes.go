package routes

import (
	. "github.com/cronjobs-service/pkg/middlewares"
	"github.com/gorilla/mux"

	. "github.com/cronjobs-service/pkg/controllers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// Home
	router.HandleFunc("/", MiddlewareJSON(GetHome)).Methods("GET")
	router.HandleFunc("/start-offer", MiddlewareJSON(StartOfferCron)).Methods("POST")
	router.HandleFunc("/stop-offer/{id}", MiddlewareJSON(StopOffer)).Methods("POST")

	return router
}
