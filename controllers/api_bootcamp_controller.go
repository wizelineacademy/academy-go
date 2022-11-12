package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"api-bootcamp/controllers/translators"
	"api-bootcamp/controllers/viewmodels"
	"api-bootcamp/mediators"
)

type ApiBootcampController struct {
	ApiBootcampFactory func() mediators.ApiMediator
}

func (ac *ApiBootcampController) GetCSVElementByID(w http.ResponseWriter, r *http.Request) {

	response := viewmodels.BaseResponse{}
	vars := mux.Vars(r)
	csvID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		response.Error = err.Error()
		SendResponse(w, http.StatusBadRequest, response)
		return
	}

	mediator := ac.ApiBootcampFactory()
	clientDto, err := mediator.GetCSVElementByID(csvID)
	if err != nil {
		fmt.Println(err)
		response.Error = err.Error()
		SendResponse(w, http.StatusBadRequest, response)
		return
	}

	response.Data = translators.ToResponseView(clientDto)
	SendResponse(w, http.StatusOK, response)
}
