package util

import (
	"Gin-Basic-Service/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var config *global.Config

func init() {
	data, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}
	global.Conf = config
}

func GetConfig() *global.Config {
	return config
}
