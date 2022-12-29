package controller

import (
	logger "github.com/georgexx009-wizeline/ondemand-go-bootcamp/utils"
	"github.com/gin-gonic/gin"
)

func (controller *Controller) ReadExternalApi(context *gin.Context) {
	pokemons, err := controller.UseCases.ConsumePokeapi()
	if err != nil {
		logger.Log(err.Error())
		context.AbortWithStatusJSON(500, err)
		return
	}

	err = controller.UseCases.SavePokemonsInCSV("pokemons", pokemons)

	context.JSON(200, pokemons)
}
