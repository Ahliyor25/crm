package handlers

import (
	"net/http"
	"strconv"

	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"github.com/gorilla/mux"
)

func (h Handler) HStatusGet(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	// Извлекаем параметр :id из URL
	vars := mux.Vars(r)
	statusID := vars["id"]

	// Преобразуем в число
	statusIDInt, err := strconv.Atoi(statusID)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}
	// Выполняем бизнес логику
	data, err := h.status.Get(uint(statusIDInt))
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = data

	return
}
