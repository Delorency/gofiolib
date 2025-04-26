package handlers

import (
	personservice "fiolib/internal/services/personService"
	"log"
	"net/http"
)

type PersonHandlerI interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Retrieve(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type personHandler struct {
	service personservice.PersonServiceI
	logger  *log.Logger
}

func NewPersonHandler(service personservice.PersonServiceI, logger *log.Logger) PersonHandlerI {
	return &personHandler{service, logger}
}
