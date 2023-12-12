package orderItem

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) Get(orderITemID uint) (data entities.OrderItem, err error) {
	data, err = p.orderItem.Get(entities.OrderItem{
		BaseGorm: entities.BaseGorm{
			ID: orderITemID,
		},
	})
	if err != nil {
		//p.logger.WithField("err", err).Error("Error while getting Product")
		return
	}
	return
}
