package config

import (
	"encoding/json"
	"sync"

	"github.com/hayrullahcansu/cachy/cross"
	"github.com/hayrullahcansu/cachy/data/configuration"
	"github.com/hayrullahcansu/cachy/framework/logging"
	"github.com/hayrullahcansu/cachy/utility"
)

var confFileName = "settings.json"
var _instance *configuration.Configuration

var _once sync.Once

func Instance() *configuration.Configuration {
	_once.Do(initialConfiguration)
	return _instance
}

func initialConfiguration() {
	_instance = initializeConfig(readConfigFromFile())
}

func initializeConfig(data string) *configuration.Configuration {
	config := configuration.Configuration{}
	json.Unmarshal([]byte(data), &config)
	return &config
}

func readConfigFromFile() string {
	b, err := utility.ReadFile(confFileName)
	if err != nil {
		logging.Infof("Cannot read config file: %s%s", confFileName, cross.NewLine)
		return ""
	}
	return string(b)
}
