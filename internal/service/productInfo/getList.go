package productinfo

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) GetList() (productInfo []entities.ProductInfo, err error) {

	return p.productInfo.GetMany()
}
