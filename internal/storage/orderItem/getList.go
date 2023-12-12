package orderItem

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"github.com/sirupsen/logrus"
)

func (p provider) GetList(target entities.OrderItem) (data []entities.OrderItem, err error) {
	err = p.postgres.Model(&entities.OrderItem{}).
		Where(&target).
		Find(&data).
		Error
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("An error occurred while retrieving OrderItem data")

		err = response.ErrInternalServer
		return
	}
	return
}
