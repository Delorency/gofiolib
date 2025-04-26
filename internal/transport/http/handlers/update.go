package handlers

import (
	"encoding/json"
	l "fiolib/internal/logger"
	"fiolib/internal/models"
	e "fiolib/internal/transport/http/error"
	response "fiolib/internal/transport/http/response"
	v "fiolib/internal/validator"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
)

type requestUpdate struct {
	Name       string `json:"name" validate:"required"`
	Surname    string `json:"surname" validate:"required"`
	Patronymic string `json:"patronymic" validate:"required"`
	Age        uint8  `json:"age" validate:"required"`
	Gender     string `json:"gender" validate:"required"`
	Nat        string `json:"nat" validate:"required"`
}

func (ph *personHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req requestUpdate

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

	bodyBytes, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		response.NewResponse(
			e.NewError(""),
			http.StatusInternalServerError,
			w,
		)
		ph.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, bodyBytes))
		return
	}

	if err := validator.New().Struct(req); err != nil {
		data, f := v.HandleValidationErrors(w, err)
		if f {
			response.NewResponse(
				data,
				http.StatusBadRequest,
				w,
			)
			ph.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, bodyBytes))
			return
		}
	}

	data := models.Person{Name: req.Name, Surname: req.Surname, Patronymic: req.Patronymic, Age: req.Age, Gender: req.Gender, Nat: req.Nat}

	obj, err := ph.service.Update(uint(id), &data)
	if err != nil {
		response.NewResponse(
			e.NewError("Ошибка обновления пользователя"),
			http.StatusInternalServerError,
			w,
		)
		ph.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, bodyBytes))
		return
	}

	response.NewResponse(
		obj,
		http.StatusOK,
		w,
	)
	ph.logger.Println(l.GetLogEntry(r, http.StatusOK, bodyBytes))
}
