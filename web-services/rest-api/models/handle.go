package models

var Config App

type App struct {
	AppConfig struct {
		Port    string `mapstructure:"port"`
		Context string `mapstructure:"context"`
	} `mapstructure:"app"`
}
