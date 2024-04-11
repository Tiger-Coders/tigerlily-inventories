package config

import (
	"flag"
	"sync"

	"github.com/Tiger-Coders/tigerlily-inventories/internal/pkg/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func LoadConfig() (config *GeneralConfig) {
	// config = configLoader()()
	/*
		💡 Setting the function to a variable allows us to call this function in a neater way. Rather than like the one above
	*/
	config = singleConfigLoader()
	return
}

var singleConfigLoader = configLoader()

func configLoader() func() *GeneralConfig {
	var appConfig *GeneralConfig
	var once sync.Once

	/*
		💡 sync.Once runs this function ONLY ONCE. This way, we don't always have to parse the config file with each calls
	*/
	return func() *GeneralConfig {
		once.Do(func() {
			var configFilePath string

			flag.StringVar(&configFilePath, "config", "config.yml", "Absolute path to configuration file")
			flag.Parse()
			appConfig = parseAndWatchConfigFile(configFilePath)
		})
		return appConfig
	}
}

func parseAndWatchConfigFile(filePath string) (config *GeneralConfig) {
	log := logger.NewLogger()
	config = &GeneralConfig{}

	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.ErrorLogger.Printf("no config file found: %+v", err)
		return &GeneralConfig{}
	}
	viperUnmarshalConfig(config, log)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.InfoLogger.Println("[CONFIG] Config has changed: ", e.Name)
		viperUnmarshalConfig(config, log)
	})

	// TODO: Refactor this. Can be neater. Goal is to perform a singleton pattern to load the config file once
	// Setting to true since we successfully oarsed the config file.
	// config.IsConfigFileProvided = true
	return
}

func viperUnmarshalConfig(config *GeneralConfig, logger *logger.Logger) {
	if err := viper.Unmarshal(config); err != nil {
		logger.ErrorLogger.Panicf("[CONFIG] Error unmarshaling app config on change : %+v\n", err)
	}
}
