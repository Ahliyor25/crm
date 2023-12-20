package order

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) Update(order entities.Order) (err error) {
	// если orderItems не пустой, то обновляем их
	if len(order.OrderItems) > 0 {
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

		// удалить c таблицы orderItems старые данные
	}

	// если status key ровняется success то вычитоваем товары

	err = p.order.Update(order)

	if err != nil {
		return
	}

	return
}
