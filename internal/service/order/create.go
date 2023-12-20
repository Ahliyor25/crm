package order

import (
	"log"

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
		productInfo, tEr := p.productInfo.Get(product.ProductInfo)
		if tEr != nil {
			err = tEr
		}
		order.OrderItems[i].ProductBasePrice = product.BasePrice - product.BasePrice*product.Discount/100
		order.OrderItems[i].ProductName = productInfo.Name

		log.Println(productInfo.Name)

		order.OrderItems[i].ProductCostPrice = product.CostPrice
		order.OrderItems[i].ProductDiscountPrice = product.Discount

		// вычисляем сумму заказа без учета скидки
		order.SubTotal += product.BasePrice * float32(item.Quantity)

		// вычисляем сумму заказа с учетом скидки
		order.Total = order.SubTotal - order.SubTotal*order.Discount/100
	}

	createdOrder, err = p.order.Create(order)
	if err != nil {
		p.logger.WithField("err", err).Error("Error while creating Order")
		return
	}

	return
}
