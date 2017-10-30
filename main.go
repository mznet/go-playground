package main

import (
	"github.com/spf13/viper"
	"log"
)

// Load AWS Configuration
func loadConfigurationFile() {
	configKeys := [4]string{"access_key", "secret", "region", "bucket"}

	viper.SetConfigFile(".aws/credential.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Check AWS Configuration nil attribute
	for _, key := range configKeys {
		if viper.Get(key) == nil {
			log.Fatalf(key + " attribute could not be set nil")
		}
	}

	log.Println("Finishig loading AWS configuration file")
}

func main() {
	log.Println("Start loading AWS configuration file")
	loadConfigurationFile()
}