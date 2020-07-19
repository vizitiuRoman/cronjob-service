package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/valyala/fasthttp"

	. "github.com/cronjob-service/pkg/routes"
	"github.com/joho/godotenv"
)

func RunCronJobsService() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error load env", err)
	}
	fmt.Println("Load env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "4041"
	}

	//ConnectDB()
	routes := InitRoutes().Handler

	fmt.Println("CronJobs-service started", port)
	log.Fatal(fasthttp.ListenAndServe(":"+port, routes))
}
