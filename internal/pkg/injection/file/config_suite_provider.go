package file

import "github.com/Tiger-Coders/tigerlily-inventories/internal/config"

func ConfigProvider() *config.GeneralConfig {
	return config.LoadConfig()
}
