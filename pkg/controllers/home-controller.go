package controllers

import (
	"fmt"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CronJobs Service Ready")
}
