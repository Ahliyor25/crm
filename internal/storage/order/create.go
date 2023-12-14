package order

import (
	"errors"
	"fmt"

	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

func (p provider) Create(target entities.Order) (order entities.Order, err error) {
	duplicateEntryError := &pgconn.PgError{Code: "23505"}

	tx := p.postgres.Begin()

	err = p.postgres.Create(&target).Error
	if err != nil {
		tx.Rollback()
		if errors.As(err, &duplicateEntryError) {
			err = response.ErrClientAlreadyRegistered
			return
		}

		p.logger.WithFields(logrus.Fields{
			"err":   err,
			"order": fmt.Sprintf("%+v", target),
		}).Error("Error while creating order")

		err = response.ErrInternalServer
		return
	}
	tx.Commit()

	order = target

	return
}
