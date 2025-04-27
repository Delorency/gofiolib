package personhandler

import (
	l "fiolib/internal/logger"
	e "fiolib/internal/transport/http/error"
	response "fiolib/internal/transport/http/response"
	sw "fiolib/internal/transport/http/swagger"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var _ = sw.SwaggerPerson{}

// @Summary Удалить пользователя
// @Accept  json
// @Produce json
// @Param   id path int true "Идентификатор пользователя"
// @Success 204 "Успешно удалено"
// @Failure 400 {object} swagger.SwaggerValidateData "Идентификатор должен быть числом"
// @Failure 404 {object} swagger.SwaggerValidateData "Объект не найден"
// @Router  /person/{id} [delete]
func (ph *personHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
		http.StatusNoContent,
		w,
	)
	ph.logger.Println(l.GetLogEntry(r, http.StatusNoContent, []byte{}))
}
