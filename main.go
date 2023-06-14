package main

import (
	"github.com/rs/zerolog/log"
	"github.com/tamnguyenwizeline/ondemand-go-bootcamp/api"

	"github.com/tamnguyenwizeline/ondemand-go-bootcamp/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	runGinServer(config)
}

func runGinServer(config util.Config) {
	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
