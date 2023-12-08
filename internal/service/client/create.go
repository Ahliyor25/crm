package client

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) Create(client entities.Client) (err error) {
	err = p.client.Create(client)
	if err != nil {
		p.logger.WithError(err).Error("Error while creating client")
		return
	}
	return
}
