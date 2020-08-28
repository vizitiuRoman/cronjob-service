package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

type CronJobServer interface {
	StartServer()
	wait(listenCh chan error, osSignals chan os.Signal)
}

type Server struct {
	Controllers fasthttp.RequestHandler
	Port        string
}

func ProvideServer() (*Server, error) {
	err := godotenv.Load()
	if err != nil {
		return &Server{}, err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4041"
	}

	return &Server{
		Controllers: initControllers().Handler,
		Port:        port,
	}, nil
}

func (srv *Server) StartServer() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	listenCh := make(chan error, 1)
	go func(listen chan error) {
		log.Println("CronJob-Service started port: " + srv.Port)
		listen <- fasthttp.ListenAndServe(":"+srv.Port, srv.Controllers)
	}(listenCh)

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	srv.wait(listenCh, osSignals)
}

func (srv *Server) wait(listenCh chan error, osSignals chan os.Signal) {
	for {
		select {
		case err := <-listenCh:
			if err != nil {
				log.Fatal("Listener error: " + err.Error())
			}
			os.Exit(0)
		case err := <-osSignals:
			log.Fatal("Shutdown signal: " + err.String())
		}
	}
}
