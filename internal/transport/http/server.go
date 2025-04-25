package http

import (
	"fiolib/internal/container"
	"fmt"
	"net/http"
)

func NewHTTPServer(addr, port string, cont *container.Container) *http.Server {
	router := NewRouter(cont)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", addr, port),
		Handler: router,
	}

	return &server
}
