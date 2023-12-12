package orderItem

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) GetList(orderId uint) (data []entities.OrderItem, err error) {
	return p.orderItem.GetList(entities.OrderItem{
		OrderID: orderId,
	})

}
