package router

import (
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/controller"
	usecases "github.com/georgexx009-wizeline/ondemand-go-bootcamp/use-cases"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) *gin.Engine {
	v1Group := r.Group("/api/v1")

	useCases := usecases.NewUseCases()
	c := controller.NewController(useCases)
	c.Router = r

	v1Group.GET("/read-external-api", c.ReadExternalApi)
	v1Group.GET("/read-pokemon-csv", c.ReadPokemonCSV)

	return r
}
