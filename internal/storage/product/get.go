package product

import (
	"fmt"

	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (p provider) Get(target entities.Product) (data entities.Product, err error) {
	err = p.postgres.Model(&entities.Product{}).
		Where(&target).
		First(&data).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = response.ErrDataNotFound
			return
		}
		p.logger.WithFields(logrus.Fields{
			"err":    err,
			"target": fmt.Sprintf("%+v", target),
		}).Error("An error occurred while retrieving Product data")

		err = response.ErrInternalServer
		return
	}

	return
}
