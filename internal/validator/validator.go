package validator

import (
	"net/http"

	e "fiolib/internal/transport/http/error"

	"github.com/go-playground/validator"
)

type validateData struct {
	e.HTTPError
	Fields []string `json:"fields"`
}

// Функция обработки ошибок валидации
func HandleValidationErrors(w http.ResponseWriter, err error) (*validateData, bool) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var vd validateData

		vd.Fields = []string{}

		for _, fieldErr := range validationErrors {
			vd.Fields = append(vd.Fields, fieldErr.Field())
		}
		vd.Err = "Не все поля заполнены"
		return &vd, true
	}

	return nil, false
}
