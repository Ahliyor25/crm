package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
)

func (h Handler) HOrderCreate(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	var data entities.Order

	// читаем orders_items и сохраняем в data
	data.OrderItems = make([]entities.OrderItem, 0)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&data)

	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}
	// сохраняем продукты в orderItems

	for i, item := range data.OrderItems {
		var product entities.Product
		product.ID = item.ProductID
		_, err := h.product.Get(product.ID)
		if err != nil {
			resp.Message = response.ErrBadRequest.Error()
			return
		}
		data.OrderItems[i].ProductBasePrice = product.BasePrice
	}

	// вычисляем общую сумму заказа

	for _, item := range data.OrderItems {
		data.Total += item.ProductBasePrice * float32(item.Quantity)
	}

	// сохраняем заказ

	err = h.order.Create(data)

	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = data

}
