package handlers

import (
	personservice "fiolib/internal/services/personService"
	"net/http"
)

type PersonHandlerI interface {
	List(w http.ResponseWriter, r *http.Request)
	// Create(w http.ResponseWriter, r *http.Request)
	// Update(w http.ResponseWriter, r *http.Request)
	// Retireve(w http.ResponseWriter, r *http.Request)
}

type personHandler struct {
	service personservice.PersonServiceI
}

func NewPersonHandler(service personservice.PersonServiceI) PersonHandlerI {
	return &personHandler{service}
}
