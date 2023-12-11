package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
)

func (h Handler) HClientGet(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	// парсим данные из запроса
	var data entities.Client

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&data)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	// выполняем бизнес  логику
	data, err = h.client.Get(data.Phone)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	// возвращаем ответ
	resp.Message = response.ErrSuccess.Error()
	resp.Payload = data
}
