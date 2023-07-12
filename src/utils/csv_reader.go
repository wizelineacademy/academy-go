package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func CsvReader(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return data, err
}
