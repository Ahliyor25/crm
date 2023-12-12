package handlers

import (
	"net/http"
	"strconv"

	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
)

func (h Handler) HProductGet(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	productID, err := strconv.Atoi(r.URL.Query().Get("product_id"))
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	// Выполняем бизнес логику
	product, err := h.product.Get(uint(productID))
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = product

}
