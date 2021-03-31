package usecases

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"

	"github.com/javargas/academy-go-q12021/entities"
	"github.com/javargas/academy-go-q12021/repository"
)

var contactList = repository.LoadData()

func HomePageHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func GetContactListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contactList)
}

func GetContactInfoHandler(w http.ResponseWriter, r *http.Request){
	keys, ok := r.URL.Query()["id"]
    
    if !ok || len(keys[0]) < 1 {
        fmt.Fprintf(w, "Url Param 'id' is missing")
        log.Println("Url Param 'id' is missing")
        return
	}
	key := keys[0]

	for _, contact := range contactList {
		if contact.Id == key {
			json.NewEncoder(w).Encode(contact)
			return
		}
	}

	json.NewEncoder(w).Encode(entities.Error{Code: 1, Message: "We could not find a fact with the specified id"})
	
	fmt.Println("Endpoint Hit: get-info")
}