package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/tamnguyenwizeline/ondemand-go-bootcamp/model"
)

func ReadCSVFile(filePath string) ([]model.Pokemon, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var result []model.Pokemon
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("invalid ID: %v", record[0])
		}

		pokemon := model.Pokemon{
			ID:   id,
			Name: record[1],
		}
		result = append(result, pokemon)
	}

	return result, nil
}
