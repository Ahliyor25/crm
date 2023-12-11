package product

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) Get(productID uint) (data entities.Product, err error) {
	data, err = p.product.Get(entities.Product{
		BaseGorm: entities.BaseGorm{
			ID: productID,
		},
	})
	if err != nil {
		//p.logger.WithField("err", err).Error("Error while getting Product")
		return
	}
	return
}
