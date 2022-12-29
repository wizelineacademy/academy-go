package usecases

import (
	"encoding/csv"
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/model"
	logger "github.com/georgexx009-wizeline/ondemand-go-bootcamp/utils"
	"os"
	"strconv"
)

func (u *UseCases) ReadPokemonCSV() ([]model.Pokemon, error) {
	file, err := os.Open("./pokemons.csv")
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	var pokemons []model.Pokemon
	for _, pokemonRecord := range records[1:] {
		pokeId, err := strconv.Atoi(pokemonRecord[0])
		pokeWeight, err := strconv.Atoi(pokemonRecord[2])
		if err != nil {
			logger.Log(err.Error())
			return nil, err
		}
		pokemons = append(pokemons, model.Pokemon{Id: pokeId, Name: pokemonRecord[1], Weight: pokeWeight})
	}

	return pokemons, nil
}
