package pkg

import (
	"fmt"
	"log"
	"net/http"
	"os"

	. "github.com/cronjobs-service/pkg/routes"
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
	routes := InitRoutes()

	fmt.Println("CronJobs-service started", port)
	log.Fatal(http.ListenAndServe(":"+port, routes))
}
