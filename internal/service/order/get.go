package order

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) Get(orderID uint) (data entities.Order, err error) {
	data, err = p.order.Get(entities.Order{
		BaseGorm: entities.BaseGorm{
			ID: orderID,
		},
	})

	if err != nil {
		//p.logger.WithField("err", err).Error("Error while getting Product")
		return
	}
	return
}
