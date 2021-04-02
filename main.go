
package main

import (
	"log"
	"net/http"
	
	"github.com/javargas/academy-go-q12021/usecases"
)

func handleRequests() {
	http.HandleFunc("/", usecases.HomePageHandler)
	http.HandleFunc("/get-jobs", usecases.GetJobListHandler)
	http.HandleFunc("/get-job-info", usecases.GetJobInfoHandler)
	http.HandleFunc("/get-jobs-api", usecases.GetJobsAPIPHandler)
	http.HandleFunc("/get-jobs-concurrent", usecases.GetJobsConcurrentPHandler)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}