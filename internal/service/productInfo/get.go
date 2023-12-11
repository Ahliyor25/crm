package productinfo

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) Get(productInfoID uint) (data entities.ProductInfo, err error) {
	data, err = p.productInfo.Get(entities.ProductInfo{
		BaseGorm: entities.BaseGorm{
			ID: productInfoID,
		},
	})
	if err != nil {
		p.logger.WithField("err", err).Error("Error while getting client")
		return
	}
	return
}
