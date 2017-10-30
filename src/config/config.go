package config

import (
	"github.com/spf13/viper"
	"logs"
)

// Load AWS Configuration
func LoadConfigurationFile() {
	logs.Info("Start loading AWS configuration file")

	configKeys := [4]string{"access_key", "secret", "region", "bucket"}

	viper.SetConfigFile(".aws/credential.yaml")

	if err := viper.ReadInConfig(); err != nil {
		logs.Error("Error reading config file, %s", err)
	}

	// Check AWS Configuration nil attribute
	for _, key := range configKeys {
		if viper.Get(key) == nil {
			logs.Error(key + " attribute could not be set nil")
		}
	}

	logs.Info("Finishig loading AWS configuration file")
}

// Set default key and value
func SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

// Set key and value
func Set(key string, value interface {}) {
	viper.Set(key, value)
}

// Get value using key
func GetString(key string) (value string){
	return viper.GetString(key)
}
