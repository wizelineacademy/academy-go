package util

import (
	"testing"

	"github.com/tamnguyenwizeline/ondemand-go-bootcamp/model"
)

func TestReadCSVFile_MultipleRecords(t *testing.T) {
	filePath := "../testdata/pokemons.csv"
	pokemon, err := ReadCSVFile(filePath)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedCount := 1281
	if len(pokemon) != expectedCount {
		t.Errorf("Expected %d pokemon, but got %d", expectedCount, len(pokemon))
	}

	expectedFirstPokemon := model.Pokemon{
		ID:   1,
		Name: "bulbasaur",
	}
	if pokemon[0] != expectedFirstPokemon {
		t.Errorf("Expected first pokemon to be %+v, but got %+v", expectedFirstPokemon, pokemon[0])
	}
}
