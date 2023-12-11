package productinfo

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
)

func (p provider) Create(productInfo entities.ProductInfo) (err error) {
	err = p.productInfo.Create(productInfo)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":    err,
			"client": productInfo,
		}).Error("Error while creating ProductInfo")
		return
	}
	return
}
