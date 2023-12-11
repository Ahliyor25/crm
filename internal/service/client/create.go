package client

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
)

func (p provider) Create(client entities.Client) (err error) {
	err = p.client.Create(client)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":    err,
			"client": client,
		}).Error("Error while creating client")
		return
	}

	return
}
