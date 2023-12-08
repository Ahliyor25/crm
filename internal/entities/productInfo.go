package entities

// ProductInfo ...
type ProductInfo struct {
	BaseGorm
	Name        string `json:"name"`
	Description string `json:"description"`
	Country     string `json:"country"`
	Code        string `json:"code"`
}
