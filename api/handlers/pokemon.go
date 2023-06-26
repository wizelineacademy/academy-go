package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hoducha/ondemand-go-bootcamp/repositories"
	pokeapi "github.com/mtslzr/pokeapi-go"
)

func GetPokemonByID(repo repositories.PokemonRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		pokemon, err := repo.GetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
			return
		}

		c.JSON(http.StatusOK, pokemon)
	}
}

func GetPokemonDetailByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		pokemon, err := pokeapi.Pokemon(strconv.Itoa(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon detail not found"})
			return
		}

		c.JSON(http.StatusOK, pokemon)
	}
}
