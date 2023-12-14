package handlers

import (
	"net/http"
	"strconv"

	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"github.com/gorilla/mux"
)

func (h Handler) HTrafficSourceGet(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	// Извлекаем параметр :id из URL
	vars := mux.Vars(r)
	trafficSourceID := vars["id"]

	// Преобразуем в число
	trafficSourceIDInt, err := strconv.Atoi(trafficSourceID)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}
	// Выполняем бизнес логику
	data, err := h.trafficSource.Get(uint(trafficSourceIDInt))
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = data

}
