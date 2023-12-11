package productpricehistory

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) Get(productInfoID uint) (data entities.ProductPriceHistory, err error) {
	data, err = p.productPriceHistory.Get(entities.ProductPriceHistory{
		BaseGorm: entities.BaseGorm{
			ID: productInfoID,
		},
	})
	if err != nil {
		p.logger.WithField("err", err).Error("Error while getting productPriceHistory data")
		return
	}
	return
}
