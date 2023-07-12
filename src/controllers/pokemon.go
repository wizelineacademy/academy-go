package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vanduc1102/ondemand-go-bootcamp/src/services"
)

func FindPokemon(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid pokemon id"})
		return
	}

	pokemon, error := services.FindById(id)
	if error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": pokemon,
	})
}
