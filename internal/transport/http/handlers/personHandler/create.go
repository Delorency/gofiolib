package personhandler

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

// @Summary Создать пользователя
// @Accept  json
// @Produce json
// @Param   person body requestCreate true "Создать пользователя"
// @Success 201 {object} models.Person
// @Failure 400 {object} v.ValidateData "Ошибка создания пользователя"
// @Failure 500 {object} e.NewError "Ошибка создания пользователя"
// @Router  /person [post]
func (ph *personHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req requestCreate

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

	obj := models.Person{Name: req.Name, Surname: req.Surname, Patronymic: req.Patronymic}

	if err := ph.service.Create(&obj); err != nil {
		response.NewResponse(
			e.NewError("Ошибка создания пользователя"),
			http.StatusInternalServerError,
			w,
		)
		ph.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, bodyBytes))
		return
	}
	response.NewResponse(
		obj,
		http.StatusCreated,
		w,
	)
	ph.logger.Println(l.GetLogEntry(r, http.StatusCreated, bodyBytes))
}
