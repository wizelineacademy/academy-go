package services

import (
	"errors"
	"os"
	"strconv"

	"github.com/vanduc1102/ondemand-go-bootcamp/src/models"
	"github.com/vanduc1102/ondemand-go-bootcamp/src/utils"
)

func FindById(id int) (models.Pokemon, error) {
	var pokemonList, error = LoadPokemonFromCSV()
	if error != nil {
		return models.Pokemon{}, error
	}
	for _, pokemon := range pokemonList {
		if pokemon.Id == id {
			return pokemon, nil
		}
	}
	return models.Pokemon{}, errors.New("pokemon id does not exist")
}

func LoadPokemonFromCSV() ([]models.Pokemon, error) {
	csvFilePath := os.Getenv("CSV_FILE")
	data, err := utils.CsvReader(csvFilePath)
	if err != nil {
		return nil, err
	}
	return createPokemonList(data)
}

func createPokemonList(data [][]string) ([]models.Pokemon, error) {
	var pokemonList []models.Pokemon

	for _, line := range data {
		var pokemon models.Pokemon
		for j, field := range line {
			if j == 0 {
				id, error := strconv.Atoi(field)
				if error != nil {
					return nil, errors.New("each item in the csv must have an id element as integer type")
				}
				pokemon.Id = id
			} else if j == 1 {
				pokemon.Name = field
			}
		}
		pokemonList = append(pokemonList, pokemon)
	}

	return pokemonList, nil
}
