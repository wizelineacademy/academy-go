package usecases

import (
	"encoding/json"
	logger "github.com/georgexx009-wizeline/ondemand-go-bootcamp/utils"
	"io/ioutil"
	"net/http"
)

var URL = "https://pokeapi.co/api/v2/pokemon/pikachu"

type PokeRes struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

func (*UseCases) ConsumePokeapi() (*PokeRes, error) {
	logger.Log("consuming poke api")
	res, err := http.Get(URL)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	byteArr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	var pokeResult PokeRes
	err = json.Unmarshal(byteArr, &pokeResult)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}

	logger.Log(pokeResult)
	return &pokeResult, nil
}
