package models

var Config App

type App struct {
	AppConig struct {
		Authorization Option `mapstructure:"authorization"`
	} `mapstructure:"app"`
}

type Option struct {
	Timeout int `mapstructure:"time-out"`
	Header  struct {
		ContentType   string `mapstructure:"content-type"`
		UserAgent     string `mapstructure:"user-agent"`
		Authorization struct {
			Type  string `mapstructure:"type"`
			Token string `mapstructure:"token"`
		} `mapstructure:"authorization"`
	} `mapstructure:"header"`
}
