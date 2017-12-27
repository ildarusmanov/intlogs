package tests

import (
	"github.com/ildarusmanov/intlogs/configs"

	"path/filepath"
)

func CreateConfig() *configs.Config {
	configFilePath, _ := filepath.Abs("./../config_test.yml")
	config := configs.LoadConfigFile(configFilePath)

	return config
}
