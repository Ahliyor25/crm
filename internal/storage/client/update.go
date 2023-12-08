package client

import (
	"fmt"

	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"github.com/sirupsen/logrus"
)

func (p provider) Update(data entities.Client) (err error) {
	err = p.postgres.Model(&entities.Client{}).
		Where(data.ID).
		Updates(&data).
		Error
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":  err,
			"data": fmt.Sprintf("%+v", data),
		}).Error("Error while updating Client data")

		err = response.ErrInternalServer
		return
	}

	return
}
