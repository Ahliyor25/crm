package entities

// Client ...
type Client struct {
	BaseGorm
	FullName string `json:"full_name"`
	Phone    string `json:"phone" validate:"min=5,max=20" gorm:"<-:create;unique"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}
