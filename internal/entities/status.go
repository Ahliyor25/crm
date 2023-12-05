package entities

// Status ...
type Status struct {
	BaseGorm
	Name string `json:"name"`
	Key  string `json:"key"`
}
