package client

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (p provider) GetList(target entities.Client) (data []entities.Client, err error) {
	err = p.postgres.Model(&entities.Client{}).
		Where(&target).
		Find(&data).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = response.ErrDataNotFound
			return
		}

		p.logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("An error occurred while retrieving Client data")

		err = response.ErrInternalServer
		return
	}

	return
}
