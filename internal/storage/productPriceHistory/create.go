package productpricehistory

import (
	"errors"

	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"github.com/jackc/pgx/v5/pgconn"
)

func (p provider) Create(productPriceHistory entities.ProductPriceHistory) (err error) {
	duplicateEntryError := &pgconn.PgError{Code: "23505"}
	tx := p.postgres.Begin()
	err = p.postgres.Create(&productPriceHistory).Error
	if err != nil {
		tx.Rollback()
		if errors.As(err, &duplicateEntryError) {
			return response.ErrorUserAlreadyExist
		}
	}
	return
}
