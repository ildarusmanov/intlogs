package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	ServerHost    string `yaml:"server_host"`
	MgoUrl        string `yaml:"mgo_url"`
	MgoDb         string `yaml:"mgo_db"`
	MgoCollection string `yaml:"mgo_collection"`
	AuthToken     string `yaml:"auth_token"`
}

func CreateNewConfig() *Config {
	return &Config{}
}

func LoadConfigFile(configFilePath string) *Config {
	configData := CreateNewConfig()

	configFileData, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(configFileData), configData)

	if err != nil {
		panic(err)
	}

	return configData
}
