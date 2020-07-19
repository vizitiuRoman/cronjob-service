package pkg

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/valyala/fasthttp"

	. "github.com/cronjob-service/pkg/routes"
	"github.com/joho/godotenv"
)

func RunCronJobsService() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error load env", err)
	}
	fmt.Println("Load env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "4041"
	}

	routes := InitRoutes().Handler
	listenCh := make(chan error, 1)
	go func(listen chan error) {
		fmt.Println("CronJob-service started", port)
		listen <- fasthttp.ListenAndServe(":"+port, routes)
	}(listenCh)

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-listenCh:
			if err != nil {
				log.Fatalf("Listener error: %s", err)
			}
			os.Exit(0)
		case err := <-osSignals:
			log.Fatalf("Shutdown signal: %s", err)
		}
	}
}
