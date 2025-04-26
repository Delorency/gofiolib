package http

import (
	"fiolib/internal/container"
	"fmt"
	"log"
	"net/http"
)

func NewHTTPServer(addr, port string, cont *container.Container, logger *log.Logger) *http.Server {
	router := NewRouter(cont, logger)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", addr, port),
		Handler: router,
	}

	return &server
}
