package usecases

import (
	"encoding/json"
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/model"
	logger "github.com/georgexx009-wizeline/ondemand-go-bootcamp/utils"
	"io/ioutil"
	"net/http"
)

var URL = "https://pokeapi.co/api/v2/pokemon/"

func getPokemon(pokemonName string) (*model.Pokemon, error) {
	logger.Log("consuming poke api")
	res, err := http.Get(URL + pokemonName)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	byteArr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	var pokeResult model.Pokemon
	err = json.Unmarshal(byteArr, &pokeResult)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	logger.Log(pokeResult)
	return &pokeResult, nil
}

func (*UseCases) ConsumePokeapi() ([]model.Pokemon, error) {
	pikachuInfo, err := getPokemon("pikachu")
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	charmanderInfo, err := getPokemon("charmander")
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	return []model.Pokemon{*pikachuInfo, *charmanderInfo}, nil
}
