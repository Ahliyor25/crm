package productpricehistory

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
)

func (p provider) Create(productPriceHistory entities.ProductPriceHistory) (err error) {
	err = p.productPriceHistory.Create(productPriceHistory)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":    err,
			"client": productPriceHistory,
		}).Error("Error while creating productPriceHistory")
		return
	}
	return
}
