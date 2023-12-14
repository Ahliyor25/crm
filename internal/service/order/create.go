package order

import (
	"github.com/ahliyor25/crm/internal/entities"
)

func (p provider) Create(order entities.Order) (createdOrder entities.Order, err error) {

	for i, item := range order.OrderItems {
		product, tErr := p.product.Get(entities.Product{
			BaseGorm: entities.BaseGorm{
				ID: item.ProductID,
			},
		})
		if tErr != nil {
			err = tErr
			return
		}
		order.OrderItems[i].ProductBasePrice = product.BasePrice
		// вычисляем сумму заказа
		order.Total += product.BasePrice * float32(item.Quantity)
	}

	createdOrder, err = p.order.Create(order)
	if err != nil {
		return
	}

	return
}
