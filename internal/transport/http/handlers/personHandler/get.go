package personhandler

import (
	l "fiolib/internal/logger"
	schemes "fiolib/internal/schemes"
	e "fiolib/internal/transport/http/error"
	response "fiolib/internal/transport/http/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// @Summary Получить список пользователей
// @Accept  json
// @Produce json
// @Param   limit query int false "Количество записей на одной странице"
// @Param   page query int false "Номер страницы"
// @Success 200 {array} models.Person
// @Failure 400 {object} e.NewError "limit должно быть числом > 0"
// @Failure 400 {object} e.NewError "page должно быть числом > 0"
// @Failure 500 {object} e.NewError "Ошибка получения данных"
// @Router  /person [get]
func (ph *personHandler) List(w http.ResponseWriter, r *http.Request) {
	var limit, page int
	var err error

	if r.URL.Query().Get("limit") == "" && r.URL.Query().Get("page") == "" {
		limit = -1
		page = 1
	} else {
		limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil || limit <= 0 {
			response.NewResponse(
				e.NewError("limit должно быть числом > 0"),
				http.StatusBadRequest,
				w,
			)
			ph.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, []byte{}))
			return
		}
		page, err = strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page <= 0 {
			response.NewResponse(
				e.NewError("page должно быть числом > 0"),
				http.StatusBadRequest,
				w,
			)
			ph.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, []byte{}))
			return
		}
	}
	pag := schemes.Pagination{Limit: limit, Page: page}
	persons, err := ph.service.List(&pag)

	if err != nil {
		response.NewResponse(
			e.NewError("Ошибка получения данных"),
			http.StatusInternalServerError,
			w,
		)
		ph.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, []byte{}))
		return
	}

	response.NewResponse(
		persons,
		http.StatusOK,
		w,
	)
	ph.logger.Println(l.GetLogEntry(r, http.StatusOK, []byte{}))
}

// @Summary Получить пользователя по id
// @Accept  json
// @Produce json
// @Param   id path int true "Идентификатор пользователя"
// @Success 200 {array} models.Person
// @Failure 400 {object} e.NewError "Идентификатор должен быть числом"
// @Failure 404 {object} e.NewError "Пользователь не найден"
// @Router  /person/{id} [get]
func (ph *personHandler) Retrieve(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		response.NewResponse(
			e.NewError("Идентификатор должен быть числом"),
			http.StatusBadRequest,
			w,
		)
		ph.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, []byte{}))
		return
	}

	group, err := ph.service.Retrieve(uint(id))

	if err != nil {
		response.NewResponse(
			e.NewError("Пользователь не найден"),
			http.StatusNotFound,
			w,
		)
		ph.logger.Println(l.GetLogEntry(r, http.StatusNotFound, []byte{}))
		return
	}

	response.NewResponse(
		group,
		http.StatusOK,
		w,
	)
	ph.logger.Println(l.GetLogEntry(r, http.StatusOK, []byte{}))
}
