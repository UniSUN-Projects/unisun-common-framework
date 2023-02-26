package models

var Config App

type App struct {
	AppConfig struct {
		Datasource Datasource `mapstructure:"datasource"`
	} `mapstructure:"app"`
}

type Datasource struct {
	Host       string     `mapstructure:"host"`
	Username   string     `mapstructure:"username"`
	Password   string     `mapstructure:"password"`
	DBName     string     `mapstructure:"db-name"`
	Port       string     `mapstructure:"port"`
	Properties Properties `mapstructure:"properties"`
}

type Properties struct {
	SSLMode  string `mapstructure:"ssl-mode"`
	TimeZone string `mapstructure:"time-zone"`
}
