package config

import (
	"fmt"
)

/*
	💡Could user wire for loggers too
*/

type GeneralConfig struct {
	PostgresConfig `mapstructure:"postgres_db" json:"postgres_db"`

	ServicePort string `mapstructure:"service_port" json:"service_port"`

	IsConfigFileProvided bool
}

type PostgresConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Name     string `mapstructure:"name" json:"name"`
	SSL      string `mapstructure:"ssl" json:"ssl"`
	MaxConn  string `mapstructure:"max_conns" json:"max_conns"`
}

func (p *PostgresConfig) GetDBString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", p.Host, p.User, p.Password, p.Name, p.Port, p.SSL)
}
