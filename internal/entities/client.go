package entities

// Client ...
type Client struct {
	BaseGorm
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}
