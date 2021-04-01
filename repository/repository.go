package repository

import (
    "log"
	"io"
	"encoding/csv"
	"os"
	
	"github.com/fatih/structs"
	"github.com/javargas/academy-go-q12021/entities"
)
const pathFile = "./DataJobs.csv"

func LoadData() []entities.Job {

	var jobList  []entities.Job = nil

	// Open the file
	csvfile, err := os.Open(pathFile)
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

		jobList = append(jobList, entities.Job {
			Uuid : record[0],
			Title : record[1],
			NormalizedJobTitle : record[2],
			ParentUuid : record[3],
		})
	}

	return jobList
}

func WriteInCSV(model entities.Job) (*os.File, error) {

	s := make([]string, 0)
	f, err := os.OpenFile(pathFile, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
        log.Fatalln("failed to open file", err)
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	for _, v := range structs.Values(model) {
		s = append(s, v.(string))
	}

	writer.Write(s)
	writer.Flush()
	return f, nil

} 