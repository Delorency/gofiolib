package handlers

import (
	l "fiolib/internal/logger"
	e "fiolib/internal/transport/http/error"
	response "fiolib/internal/transport/http/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (ph *personHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		response.NewResponse(
			e.NewError("Идентификатор должен быть числом"),
			http.StatusBadGateway,
			w,
		)
		ph.logger.Println(l.GetLogEntry(r, http.StatusBadGateway, []byte{}))
		return
	}
	err = ph.service.Delete(uint(id))

	if err != nil {
		fmt.Println(err)
		response.NewResponse(
			e.NewError("Объект не найден"),
			http.StatusNotFound,
			w,
		)
		ph.logger.Println(l.GetLogEntry(r, http.StatusNotFound, []byte{}))
		return
	}

	response.NewResponse(
		[]byte{},
		http.StatusOK,
		w,
	)
	ph.logger.Println(l.GetLogEntry(r, http.StatusOK, []byte{}))
}
