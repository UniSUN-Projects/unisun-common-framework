package models

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
