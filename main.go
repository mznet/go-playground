package main

import (
	"logs"
	"config"
)

func main() {
	config.LoggerSetting()
	config.LoadConfigurationFile()
	logs.Initialize()


	logs.Info("test")

}