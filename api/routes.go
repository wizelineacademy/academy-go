package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"api-bootcamp/controllers"
	"api-bootcamp/mediators"
)

func Routes() http.Handler {

	apiBootcampController := generateController()

	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	router.HandleFunc("/{id:[0-9]+}", apiBootcampController.GetCSVElementByID).Methods(http.MethodGet)

	return router
}

func generateController() controllers.ApiBootcampController {

	apiBootcampFactory := func() mediators.ApiMediator {
		return mediators.NewApiMediator()
	}
	return controllers.ApiBootcampController{
		ApiBootcampFactory: apiBootcampFactory,
	}
}
