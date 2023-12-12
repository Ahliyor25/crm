package trafficsource

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/sirupsen/logrus"
)

func (p provider) Create(trafficSource entities.TrafficSource) (err error) {
	err = p.trafficSource.Create(trafficSource)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":    err,
			"status": trafficSource,
		}).Error("Error while creating trafficSource")
		return
	}

	return
}
