package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataFile = "pokemon_data.csv"

func TestGetPokemonByID(t *testing.T) {
	router := setupRouter(testDataFile)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/pokemon/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"id\":1,\"name\":\"bulbasaur\"}", w.Body.String())
}

func TestGetPokemonByInvalidID(t *testing.T) {
	router := setupRouter(testDataFile)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/pokemon/invalid", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{\"error\":\"Invalid ID\"}", w.Body.String())
}

func TestGetPokemonByIDNotFound(t *testing.T) {
	router := setupRouter(testDataFile)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/pokemon/1000", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "{\"error\":\"Pokemon not found\"}", w.Body.String())
}