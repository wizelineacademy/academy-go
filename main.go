package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vanduc1102/ondemand-go-bootcamp/src/controllers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP",
			"at":     time.Now(),
		})
	})

	router.GET("pokemon/:id", controllers.FindPokemon)
	router.Run("localhost:8080")
}
