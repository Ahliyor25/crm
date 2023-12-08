package client

import (
	"fmt"

	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (p provider) GetMany(target entities.Client, currentClientID uint) (clients []entities.Client, err error) {
	// Use current Client ID to filter out the current Client
	err = p.postgres.Model(&entities.Client{}).
		Where("id != ?", currentClientID).
		Where(&target).
		Find(&clients).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = response.ErrDataNotFound
			return
		}

		p.logger.WithFields(logrus.Fields{
			"err":    err,
			"target": fmt.Sprintf("%+v", target),
		}).Error("An error occurred while retrieving client data")

		err = response.ErrInternalServer
		return
	}

	return
}
