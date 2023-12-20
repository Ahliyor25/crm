package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahliyor25/crm/internal/entities"
	"github.com/ahliyor25/crm/pkg/bootstrap/http/misc/response"
)

func (h Handler) HAuthLogin(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	var data entities.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}
	token, refreshToken, err := h.auth.Login(data.Email, data.Password)

	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	resp.Message = response.ErrSuccess.Error()
	resp.Payload = map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
	}
}

// swagger:operation POST /auth/login Authorization Authorization
//
// Авторизация
//
// ## Роут предназначен для авторизации пользователя в системе
//
// ---
//
// responses:
//
//	200:
//	  description: Pong ! 🥎
//	  schema:
//	    "$ref": "#/responses/success/schema"
