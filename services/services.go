package services

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/javargas/academy-go-q12021/entities"
	"github.com/javargas/academy-go-q12021/repository"
)

func GetJobsAPI() (entities.Job, *entities.Error){

	var responseObject entities.Job
	var error entities.Error
	
    response, err := http.Get("http://api.dataatwork.org/v1/jobs/26bc4486dfd0f60b3bb0d8d64e001800")
    if err != nil {
		log.Fatal(err)
		error.Message = err.Error()
		return responseObject, &error
	}	
	

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
		error.Message = err.Error()
		return responseObject, &error
    }

    json.Unmarshal(responseData, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)

	repository.WriteInCSV(responseObject)

	return responseObject, nil
}