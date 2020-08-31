package main

import (
	"github.com/cronjob-service/pkg/server"
)

func main() {
	cronJobService := server.NewService()
	cronJobService.StartService()
}
