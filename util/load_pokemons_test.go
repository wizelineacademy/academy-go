package util

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"
)

func TestLoadPokemonData_ValidURL(t *testing.T) {
	filePath := "../testdata/pokemons.csv"

	err := LoadPokemonData(filePath)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify that the CSV file is created and contains the expected data
	file, err := os.Open(filePath)
	if err != nil {
		t.Errorf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Errorf("Error reading CSV file: %v", err)
	}

	expectedCount := 1281
	if len(records) != expectedCount {
		t.Errorf("Expected %d records, but got %d", expectedCount, len(records))
	}

	expectedFirstRecord := []string{"1", "bulbasaur"}
	if !reflect.DeepEqual(records[0], expectedFirstRecord) {
		t.Errorf("Expected first record to be %+v, but got %+v", expectedFirstRecord, records[0])
	}
}
