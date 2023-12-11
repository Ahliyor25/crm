package client

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) Get(phone string) (data entities.Client, err error) {
	data, err = p.client.Get(entities.Client{
		Phone: phone,
	})
	if err != nil {
		p.logger.WithField("err", err).Error("Error while getting client")
		return
	}
	return

}
