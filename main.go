
package main

import (
	"log"
	"net/http"
	
	"github.com/javargas/academy-go-q12021/usecases"
)

func handleRequests() {
	http.HandleFunc("/", usecases.HomePageHandler)
	http.HandleFunc("/get-contacts", usecases.GetContactListHandler)
	http.HandleFunc("/get-contact-info", usecases.GetContactInfoHandler)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}