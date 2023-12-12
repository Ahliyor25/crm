package product

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
)

func (p provider) Create(product entities.Product) (err error) {
	err = p.product.Create(product)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":     err,
			"product": product,
		}).Error("Error while creating Product")
		return
	}
	return
}
