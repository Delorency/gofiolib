package http

import (
	"fiolib/internal/container"
	"log"

	_ "fiolib/docs"
	hP "fiolib/internal/transport/http/handlers/personHandler"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func AddMiddleware(router *chi.Mux) *chi.Mux {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	return router
}

func NewRouter(cont *container.Container, logger *log.Logger) *chi.Mux {
	router := AddMiddleware(chi.NewRouter())
	handlers := hP.NewPersonHandler(cont.PersonService, logger)

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	router.Post("/person", handlers.Create)
	router.Get("/person", handlers.List)
	router.Get("/person/{id}", handlers.Retrieve)
	router.Put("/person/{id}", handlers.Update)
	router.Delete("/person/{id}", handlers.Delete)

	return router
}
