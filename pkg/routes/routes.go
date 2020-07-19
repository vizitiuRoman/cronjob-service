package routes

import (
	. "github.com/cronjob-service/pkg/controllers"
	. "github.com/fasthttp/router"
)

func InitRoutes() *Router {
	router := New()

	router.GET("/api/jobs", GetJobs)
	router.POST("/api/offer", StartOfferJob)
	router.DELETE("/api/offer/{id}", DeleteOfferJob)

	return router
}
