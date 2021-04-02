package services

import (
	"fmt"
	"os"
	"log"
	"math"
	"sync"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/javargas/academy-go-q12021/entities"
	"github.com/javargas/academy-go-q12021/repository"
)

func GetJobsAPI() (entities.Job, *entities.Error){
	
    response, err := http.Get("http://api.dataatwork.org/v1/jobs/26bc4486dfd0f60b3bb0d8d64e001800")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
	}	
	

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject entities.Job
    json.Unmarshal(responseData, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)

	repository.WriteInCSV(responseObject)

	return responseObject, nil
}

func calculatePoolSize(items int, itemsPerWorker int, totalJobs int) int {
	var poolSize int
	if items%itemsPerWorker != 0 {
		poolSize = int(math.Ceil(float64(items) / float64(itemsPerWorker)))
	} else {
		poolSize = int(items / itemsPerWorker)
	}

	// If we overpass the number of workers above the half of number
	// of items it's gonna get into an infinit looop
	if poolSize > (totalJobs / 2) {
		poolSize = totalJobs / 2
	}
	return poolSize
} 

func calculateMaxJobs(totalJobs int) int {
	var maxPokemons int

	if totalJobs%2 == 0 {
		maxPokemons = totalJobs / 2
	} else {
		maxPokemons = totalJobs/2 + 1
	}
	return maxPokemons
}

func GetJobsConcurrently(typeNumber string, items int, itemsPerWorker int) ([]entities.Job, *entities.Error) {
	
	var jobList = repository.LoadData()

	filterType := entities.TypeNumberFilter{
		Even: "even",
		Odd:  "odd",
	}
	totalJobs:= len(jobList)
	poolSize := calculatePoolSize(items, itemsPerWorker, totalJobs)
	maxJobs := calculateMaxJobs(totalJobs)

	values := make(chan int)
	tasks := make(chan int, poolSize)
	shutdown := make(chan struct{})

	startIndex := 0
	var limit int
	limit = int(math.Ceil(float64(totalJobs) / float64(poolSize)))
	lastLimit := (totalJobs % limit)

	var wg sync.WaitGroup
	wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		go func(jobs <-chan int) {
			for {
				var id string
				var limitRecalculated int
				start := <-jobs

				// We do need to iterate with the same limit every time.
				// on the last cycle we use the leftovers of the division (modulus)
				if limit+start >= totalJobs && lastLimit != 0 { // lastLimit can be 0, take care of that
					limitRecalculated = start + lastLimit
				} else {
					limitRecalculated = start + limit
				}

				for j := start; j < limitRecalculated; j++ {
					id = jobList[j].Uuid

					select {
					case values <- id:
					case <-shutdown:
						wg.Done()
						return
					}
				}
			}
		}(tasks)
	}

	for i := 0; i < poolSize; i++ {
		tasks <- startIndex
		startIndex += limit
	}
	close(tasks)

	var filteredPokemons []entities.Job = nil
	bucket := make(map[int]int, totalJobs+1)
	for elem := range values {
		if typeNumber == filterType.Odd {
			if elem%2 != 0 && bucket[elem] == 0 {
				filteredPokemons = append(filteredPokemons, pokemons[elem-1])
				bucket[elem] = elem // we use the map to mark the ones that has been added to the collection
			}
		} else if typeNumber == filterType.Even {
			if elem%2 == 0 && bucket[elem] == 0 {
				filteredPokemons = append(filteredPokemons, pokemons[elem-1])
				bucket[elem] = elem // we use the map to mark the ones that has been added to the collection
			}
		}
		if len(filteredPokemons) >= items || len(filteredPokemons) >= maxJobs {
			break // Finally if we reahc the items value or the possibly half that we cna take, break the loop
		}
	}

	// closing this channel we send the signal to all the goroutines to be finished
	close(shutdown)

	wg.Wait()

	return filteredPokemons, nil
}