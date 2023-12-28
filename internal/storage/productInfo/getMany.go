package productinfo

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (p provider) GetMany() (productInfo []entities.ProductInfo, err error) {
	// Use current user ID to filter out the current user

	err = p.postgres.Model(&entities.ProductInfo{}).
		Find(&productInfo).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = response.ErrDataNotFound
			return
		}
		p.logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("An error occurred while retrieving productInfo data")

		err = response.ErrInternalServer
		return
	}

	return
}
