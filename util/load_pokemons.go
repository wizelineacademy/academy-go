package util

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/tamnguyenwizeline/ondemand-go-bootcamp/model"
)

type PokemonResponse struct {
	model.Pokemon
	URL string
}

func LoadPokemonData(filePath string) error {
	url := "https://pokeapi.co/api/v2/pokemon?limit=100000&offset=0"
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create a new CSV file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Read the API response and write data to CSV
	decoder := json.NewDecoder(response.Body)
	var data struct {
		Results []PokemonResponse `json:"results"`
	}
	err = decoder.Decode(&data)
	if err != nil {
		return err
	}

	// Extract ID from the URL and write each row to CSV
	for _, pokemon := range data.Results {
		id := extractIDFromURL(pokemon.URL)
		err = writer.Write([]string{strconv.Itoa(id), pokemon.Name})
		if err != nil {
			return err
		}
	}

	fmt.Println("Data saved to", filePath)
	return nil
}

func extractIDFromURL(url string) int {
	// Example URL: "https://pokeapi.co/api/v2/pokemon/25/"
	idStr := url[len("https://pokeapi.co/api/v2/pokemon/"):]
	idStr = idStr[:len(idStr)-1]
	id, _ := strconv.Atoi(idStr)
	return id
}
