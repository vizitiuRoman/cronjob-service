package server

import (
	. "github.com/cronjob-service/pkg/controllers"
	. "github.com/fasthttp/router"
)

func initControllers() *Router {
	router := New()

	router.GET("/api/jobs", GetJobs)
	router.POST("/api/offer", StartOfferJob)
	router.PUT("/api/offer", UpdateOfferJob)
	router.DELETE("/api/offer/{id}", DeleteOfferJob)

	return router
}
