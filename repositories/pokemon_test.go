package repositories

import (
	"errors"
	"strings"
	"testing"

	"github.com/hoducha/ondemand-go-bootcamp/models"
)

func TestCSVRepository_GetByID(t *testing.T) {
	pokemonData := map[int]*models.Pokemon{
		1: {ID: 1, Name: "Bulbasaur"},
		2: {ID: 2, Name: "Charmander"},
	}

	repo := &CSVRepository{
		pokemonData: pokemonData,
	}

	// Test case: Existing Pokemon
	pokemon, err := repo.GetByID(1)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if pokemon == nil {
		t.Error("Expected Pokemon, got nil")
	} else if pokemon.ID != 1 || pokemon.Name != "Bulbasaur" {
		t.Errorf("Expected Pokemon with ID 1 and name 'Bulbasaur', got: %+v", pokemon)
	}

	// Test case: Non-existing Pokemon
	pokemon, err = repo.GetByID(3)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if pokemon != nil {
		t.Errorf("Expected nil Pokemon, got: %+v", pokemon)
	}
	expectError := errors.New("Pokemon not found")
	if err.Error() != expectError.Error() {
		t.Errorf("Expected '%v', got: %v", expectError, err)
	}

}

func TestNewPokemonRepository(t *testing.T) {
	// Mock CSV data
	csvData := `1,Bulbasaur
2,Charmander`
	r := strings.NewReader(csvData)

	// Test case: Valid CSV file
	repo, err := NewPokemonRepositoryFromReader(r)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if repo == nil {
		t.Error("Expected repository instance, got nil")
	}

	// Test case: Invalid CSV file
	invalidCSV := "invalid csv data"
	r = strings.NewReader(invalidCSV)
	_, err = NewPokemonRepositoryFromReader(r)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
