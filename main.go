package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hoducha/ondemand-go-bootcamp/api"
	"github.com/hoducha/ondemand-go-bootcamp/repositories"
)

func setupRouter(dataFile string) *gin.Engine {
	pokemonRepo, err := repositories.NewPokemonRepository(dataFile)
	if err != nil {
		log.Fatalf("Failed to initialize Pokemon repository: %v", err)
	}

	router := gin.Default()
	api.SetupRoutes(router, pokemonRepo)

	return router
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <pokemon.csv>", os.Args[0])
	}
	dataFile := os.Args[1]
		
	router := setupRouter(dataFile)

	log.Fatal(router.Run(":8080"))
}
