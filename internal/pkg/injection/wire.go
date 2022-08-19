//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/Tiger-Coders/tigerlily-inventories/internal/config"

	"github.com/Tiger-Coders/tigerlily-inventories/internal/pkg/injection/file"
	"github.com/google/wire"
)

func GetAppConfig() *config.GeneralConfig {
	wire.Build(file.ConfigProvider)
	return &config.GeneralConfig{}
}

func GetDBString() string {
	wire.Build(file.DBStringProvider)
	return GetAppConfig().GetDBString()
}
