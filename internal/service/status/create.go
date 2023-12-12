package status

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
)

func (p provider) Create(status entities.Status) (err error) {
	err = p.status.Create(status)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":    err,
			"status": status,
		}).Error("Error while creating client")
		return
	}

	return
}
