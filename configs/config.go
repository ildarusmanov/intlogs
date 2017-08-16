package configs

import (
    "fmt"
    "io/ioutil"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

type Config struct {
	ServerHost string `yaml:"server_host"`
	MgoUrl     string `yaml:"mgo_url"`
	MgoDb      string `yaml:"mgo_db"`
	AuthToken  string `yaml:"auth_koen"`
}

func CreateNewConfig() *Config {
	return &Config{}
}

func LoadConfigFile() *Config {
	configData := CreateNewConfig()

	configFilePath, _ := filepath.Abs("./config.yml")
	configFileData, err := ioutil.ReadFile(configFilePath)

	if (err != nil) {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(configFileData), configData)

	if (err != nil) {
		panic(err)
	}

	return configData
}