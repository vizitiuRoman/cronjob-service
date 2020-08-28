package main

import (
	"fmt"
	"os"

	"github.com/cronjob-service/pkg/server"
)

func main() {
	service, err := server.ProvideServer()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
		return
	}
	service.StartServer()
}
