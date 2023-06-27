package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tamnguyenwizeline/ondemand-go-bootcamp/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config util.Config
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config) (*Server, error) {

	server := &Server{
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/pokemons", server.getPokemons)
	router.POST("/pokemons", server.loadPokemons)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
