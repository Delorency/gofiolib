package http

import (
	"fiolib/internal/container"

	hP "fiolib/internal/transport/http/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func AddMiddleware(router *chi.Mux) *chi.Mux {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	return router
}

func NewRouter(cont *container.Container) *chi.Mux {
	router := AddMiddleware(chi.NewRouter())
	handlers := hP.NewPersonHandler(cont.PersonService)

	router.Get("/", handlers.List)

	return router
}
