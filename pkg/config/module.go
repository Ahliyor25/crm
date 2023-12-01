package config

import (
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(ImportConfigs)

// Postgres configuration
type Postgres struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         uint16 `json:"port"`
	SSLMode      string `mapstructure:"ssl_mode"`
	DatabaseName string `mapstructure:"database_name"`
}

// Server configuration
type Server struct {
	Host      string `json:"host"`
	Port      uint16 `json:"port"`
	SecretKey string `mapstructure:"secret_key"`
}

// Config — struct is designed to combine the
// structures used
// 
// Add or remove structures that are contained
// in your `config.json` file to this structure
type Config struct {
	Postgres Postgres `json:"postgres"`
	Server   Server   `json:"server"`
}
