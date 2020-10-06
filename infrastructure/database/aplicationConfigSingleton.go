package database

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// AppConfig is a singleton to get configuration
type AppConfig struct {
	Database *Database `json:"database,omitempty"`
}

// GetManagerAppConfig get singleton of  aplications config
func GetManagerAppConfig() *AppConfig {
	var once sync.Once
	var singleton *AppConfig
	once.Do(func() {
		singleton = NewAppConfig()
	})
	return singleton
}

// LoadConfiguration load configurations file
func (a *AppConfig) LoadConfiguration(file string) AppConfig {
	var config AppConfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config

}

// NewAppConfig Create Aplication config
func NewAppConfig() *AppConfig {
	appConfig := AppConfig{}
	appConfig.LoadConfiguration("appconfig.json")
	return &appConfig
}
