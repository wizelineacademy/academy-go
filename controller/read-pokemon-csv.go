package controller

import (
	logger "github.com/georgexx009-wizeline/ondemand-go-bootcamp/utils"
	"github.com/gin-gonic/gin"
)

func (controller *Controller) ReadPokemonCSV(context *gin.Context) {
	pokemons, err := controller.UseCases.ReadPokemonCSV()
	if err != nil {
		logger.Log(err.Error())
		context.AbortWithStatusJSON(500, err)
		return
	}

	context.JSON(200, pokemons)
}
