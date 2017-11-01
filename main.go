package main

import (
	"logs"
	"config"
)

func main() {
	config.LoadConfigurationFile()

	logs.Info("test")

}