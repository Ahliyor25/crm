package auth

import (
	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/utils"
)

func (p provider) Login(email, password string) (token, refreshToken string, err error) {

	user, err := p.user.Get(entities.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return
	}
	token, refreshToken, err = utils.CreateToken(user.ID)
	if err != nil {
		p.logger.WithError(err).Error("Error while creating token")
		return
	}

	return
}
