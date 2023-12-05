package entities

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Validator; tag required_if is used as an optional field
// Use a single instance of Validate, it caches struct info
var validate *validator.Validate

// BaseGorm — basic structure of the gorm package with the
// adjustments made in the tags
//
// swagger:ignore
type BaseGorm struct {
	// required: true
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// BaseGormWithUUID — basic structure of the gorm package with the
// adjustments made in the tags
//
// swagger:ignore
type BaseGormWithUUID struct {
	// required: true
	ID        uuid.UUID      `json:"id" gorm:"<-:create;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// BaseKeycloak — the basic structure responsible for
// obtaining authorization data when a user logs in
type BaseKeycloak struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
