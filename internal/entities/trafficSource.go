package entities

// TrafficSource ...
type TrafficSource struct {
	BaseGorm
	Name string `json:"name"`
	Link string `json:"link"`
}
