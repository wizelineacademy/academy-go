package main

import (
	"api-bootcamp/api"
	"net/http"
)

func main() {

	router := api.Routes()

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	srv.ListenAndServe()
}
