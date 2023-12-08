package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
)

func (h Handler) HClientCreate(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	// парсим данные из запроса
	var data entities.Client

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// проверяем что все поля пришли
	err := decoder.Decode(&data)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	// save client
	err = h.client.Create(data)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}
}
