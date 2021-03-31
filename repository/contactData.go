package repository

import (
    "log"
	"io"
	"encoding/csv"
	"os"

	"github.com/javargas/academy-go-q12021/entities"
)

func LoadData() []entities.Contact {

	var contactList  []entities.Contact = nil

	// Open the file
	csvfile, err := os.Open("DataContacts.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	reader := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		contactList = append(contactList, entities.Contact {
			Id : record[0],
			Nombre: record[1],
			PhoneNumber: record[2],
		})
	}

	return contactList
}