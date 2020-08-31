package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type CronJobService interface {
	StartService()
	wait()
}

type Service struct {
	controllers fasthttp.RequestHandler
	port        string
	osSignals   chan os.Signal
	listenCh    chan error
}

func init() {
	err := godotenv.Load()
	if err != nil {
		zap.S().Fatalf("Load env error: %v", err)
	}

	err = initLogger()
	if err != nil {
		zap.S().Fatalf("InitLogger error: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		zap.S().Fatalf("PORT env does not exist")
	}
}

func NewService() *Service {
	return &Service{
		controllers: initControllers().Handler,
		port:        os.Getenv("PORT"),
		osSignals:   make(chan os.Signal, 1),
		listenCh:    make(chan error, 1),
	}
}

func (srv *Service) StartService() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	go func(listen chan error) {
		zap.S().Info("CronJob Service started on port: " + srv.port)
		listen <- fasthttp.ListenAndServe(":"+srv.port, srv.controllers)
	}(srv.listenCh)

	signal.Notify(srv.osSignals, syscall.SIGINT, syscall.SIGTERM)

	srv.wait()
}

func (srv *Service) wait() {
	for {
		select {
		case err := <-srv.listenCh:
			if err != nil {
				zap.S().Fatalf("Listener error: %v", err)
			}
			os.Exit(0)
		case err := <-srv.osSignals:
			zap.S().Fatalf("Shutdown signal: %v", err.String())
		}
	}
}
