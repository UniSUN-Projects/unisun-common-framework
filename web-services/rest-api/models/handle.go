package models

var Config App

type App struct {
	AppConfig struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"app"`
}
