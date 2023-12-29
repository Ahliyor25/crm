package handlers

import (
	"net/http"

	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
)

func (h Handler) HProductGetList(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	// Выполняем бизнес логику
	product, err := h.product.GetList()
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = product

}
