package config

import (
	"fmt"
)

/*
	💡Could user wire for loggers too
*/

type GeneralConfig struct {
	PostgresHost    string `mapstructure:"postgres_host" json:"postgres_host"`
	PostgresUser    string `mapstructure:"postgres_user"`
	PostgresDBName  string `mapstructure:"postgres_db_name" json:"postgres_db_name"`
	PostgresSSLMode string `mapstructure:"postgres_ssl_mode" json:"postgres_ssl_mode"`
	PostgresPort    string `mapstructure:"postgres_port" json:"postgres_port"`

	ServicePort string `mapstructure:"service_port" json:"service_port"`

	IsConfigFileProvided bool
}

func (c *GeneralConfig) GetDBString() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s", c.PostgresHost, c.PostgresUser, c.PostgresDBName, c.PostgresPort, c.PostgresSSLMode)
}
