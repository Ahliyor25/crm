package order

import (
	"fmt"

	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"github.com/sirupsen/logrus"
)

func (p provider) Update(target entities.Order) (err error) {
	err = p.postgres.Model(&entities.Order{}).
		Where("id = ?", target.ID).
		Updates(&target).
		Error
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":  err,
			"data": fmt.Sprintf("%+v", target),
		}).Error("Error while updating order data")

		err = response.ErrInternalServer
		return
	}
	return
}
