package controllers

import (
	"api-bootcamp/controllers/viewmodels"
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, httpStatus int, data viewmodels.BaseResponse) {
	w.Header().Add("Content-Type", "application/json")
	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(httpStatus)
	w.Write(bytes)
}
