package usecases

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"

	"github.com/javargas/academy-go-q12021/entities"
	"github.com/javargas/academy-go-q12021/repository"
	"github.com/javargas/academy-go-q12021/services"
)

var jobList = repository.LoadData()

func HomePageHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func GetJobListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobList)
}

func GetJobInfoHandler(w http.ResponseWriter, r *http.Request){
	keys, ok := r.URL.Query()["id"]
    
    if !ok || len(keys[0]) < 1 {
        fmt.Fprintf(w, "Url Param 'id' is missing")
        log.Println("Url Param 'id' is missing")
        return
	}
	key := keys[0]

	for _, job := range jobList {
		if job.Uuid == key {
			json.NewEncoder(w).Encode(job)
			return
		}
	}

	json.NewEncoder(w).Encode(entities.Error{Code: 1, Message: "We could not find a job with the specified id"})
	
	fmt.Println("Endpoint Hit: get-info")
}

func GetJobsAPIPHandler(w http.ResponseWriter, r *http.Request) {
	
	var job, err = services.GetJobsAPI()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(job)
	}
}