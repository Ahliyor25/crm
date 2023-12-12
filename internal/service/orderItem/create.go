package orderItem

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
)

func (p provider) Create(orderItem entities.OrderItem) (err error) {
	err = p.orderItem.Create(orderItem)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":     err,
			"product": orderItem,
		}).Error("Error while creating orderItem")
		return
	}
	return
}
