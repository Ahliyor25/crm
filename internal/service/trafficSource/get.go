package trafficsource

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) Get(trafficSourceID uint) (data entities.TrafficSource, err error) {
	data, err = p.trafficSource.Get(entities.TrafficSource{
		BaseGorm: entities.BaseGorm{
			ID: trafficSourceID,
		},
	})
	if err != nil {
		p.logger.WithField("err", err).Error("Error while getting status")
		return
	}
	return

}
