package product

import "github.com/ahliyor25/crm/internal/entities"

func (p provider) GetList() (product []entities.Product, err error) {
	return p.product.GetMany()
}
