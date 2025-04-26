package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	l "fiolib/internal/logger"
	"fiolib/internal/models"
	e "fiolib/internal/transport/http/error"
	response "fiolib/internal/transport/http/response"
	v "fiolib/internal/validator"

	"github.com/go-playground/validator"
)

type requestCreate struct {
	Name       string `json:"name" validate:"required"`
	Surname    string `json:"surname" validate:"required"`
	Patronymic string `json:"patronymic" validate:"required"`
}

func (gh *personHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req requestCreate

	bodyBytes, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		response.NewResponse(
			e.NewError(""),
			http.StatusInternalServerError,
			w,
		)
		gh.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, bodyBytes))
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
			gh.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, bodyBytes))
			return
		}
	}

	obj := models.Person{Name: req.Name, Surname: req.Surname, Patronymic: req.Patronymic}

	if err := gh.service.Create(&obj); err != nil {
		response.NewResponse(
			e.NewError("Ошибка создания пользователя"),
			http.StatusInternalServerError,
			w,
		)
		gh.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, bodyBytes))
		return
	}
	response.NewResponse(
		obj,
		http.StatusCreated,
		w,
	)
	gh.logger.Println(l.GetLogEntry(r, http.StatusCreated, bodyBytes))
}
