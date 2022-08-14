package config

import (
	"github.com/Tiger-Coders/tigerlily-inventories/internal/pkg/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type GeneralConfig struct {
	PostgresDB string `mapstructure:"postgresDB"`

	IsDBWithEnv bool
}

func LoadConfig() (config *GeneralConfig) {
	log := logger.NewLogger()

	config = &GeneralConfig{}
	viper.SetConfigFile("./config.yml")
	viper.ReadInConfig()

	if err := viper.Unmarshal(config); err != nil {
		log.ErrorLogger.Fatalf("Error reading/finding config file: %+v", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.InfoLogger.Println("[CONFIG] Inventory Config has changed: ", e.Name)
		if err := viper.Unmarshal(config); err != nil {
			log.ErrorLogger.Panicf("[CONFIG] Error unmarshaling app config on change : %+v\n", err)
		}
	})
	return
}
