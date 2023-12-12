package status

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) Get(statusID uint) (data entities.Status, err error) {
	data, err = p.status.Get(entities.Status{
		BaseGorm: entities.BaseGorm{
			ID: statusID,
		},
	})
	if err != nil {
		p.logger.WithField("err", err).Error("Error while getting status")
		return
	}
	return

}
