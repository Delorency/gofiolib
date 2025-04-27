package util

import (
	schemes "fiolib/internal/schemes"
	"net/http"
	"strconv"
)

func GetFilter(r *http.Request) *schemes.Filter {
	// Извлекаем параметры из query string
	age, err := strconv.Atoi(r.URL.Query().Get("age"))
	if err != nil {
		age = 0
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		id = 0
	}

	return &schemes.Filter{
		ID:         uint(id),
		Name:       r.URL.Query().Get("name"),
		Surname:    r.URL.Query().Get("surname"),
		Patronymic: r.URL.Query().Get("patronymic"),
		Age:        uint8(age),
		Gender:     r.URL.Query().Get("gender"),
		Nat:        r.URL.Query().Get("nat"),
		CreatedAt:  r.URL.Query().Get("created_at"),
		UpdatedAt:  r.URL.Query().Get("updated_at"),
		DeletedAt:  r.URL.Query().Get("deleted_at"),
	}
}
