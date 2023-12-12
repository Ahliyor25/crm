package order

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
)

func (p provider) Create(order entities.Order) (err error) {
	err = p.order.Create(order)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":     err,
			"product": order,
		}).Error("Error while creating Order")
		return
	}
	return
}
