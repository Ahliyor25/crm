package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
)

func (h Handler) HStatusGet(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	var data entities.Status
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&data)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = data

	return
}
