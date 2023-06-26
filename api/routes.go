package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hoducha/ondemand-go-bootcamp/api/handlers"
	"github.com/hoducha/ondemand-go-bootcamp/repositories"
)

func SetupRoutes(router *gin.Engine, repo repositories.PokemonRepository) {

	v1 := router.Group("/v1")
	{
		v1.GET("/pokemon/:id", handlers.GetPokemonByID(repo))
		v1.GET("/pokemon/:id/detail", handlers.GetPokemonDetailByID())
	}
}
