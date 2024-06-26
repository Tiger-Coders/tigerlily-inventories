// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/Tiger-Coders/tigerlily-inventories/internal/config"
	"github.com/Tiger-Coders/tigerlily-inventories/internal/pkg/injection/file"
)

// Injectors from wire.go:

func GetAppConfig() *config.GeneralConfig {
	generalConfig := file.ConfigProvider()
	return generalConfig
}

func GetDBString() string {
	string2 := file.DBStringProvider()
	return string2
}
