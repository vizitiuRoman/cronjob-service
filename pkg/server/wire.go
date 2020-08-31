package server

import (
	"os"

	. "github.com/cronjob-service/pkg/controllers"
	. "github.com/fasthttp/router"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLogger() error {
	var config zap.Config
	if os.Getenv("LOG_LEVEL") == "prod" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()
	zap.ReplaceGlobals(logger)
	if err != nil {
		return err
	}
	return nil
}

func initControllers() *Router {
	router := New()

	router.GET("/api/jobs", GetJobs)
	router.POST("/api/offer", StartOfferJob)
	router.PUT("/api/offer", UpdateOfferJob)
	router.DELETE("/api/offer/{id}", DeleteOfferJob)

	return router
}
