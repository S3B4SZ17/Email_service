package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	// mapstructure is used since is what Viper uses under the hood
	Smtp_server Smtp_server `mapstructure:"smtp_server"`
	Http_server Http_server `mapstructure:"http_server"`
}

type Http_server struct {
	Gin_mode string `mapstructure:"gin_mode"`
	HttpPort string `mapstructure:"http_port"`
	Cors     struct {
		Enabled    bool     `mapstructure:"enabled"`
		List_hosts []string `mapstructure:"list_hosts"`
	} `mapstructure:"cors"`
}

type Smtp_server struct {
	Email_from string `mapstructure:"emailFrom"`
	Host_url   string `mapstructure:"host_url"`
	Port       int    `mapstructure:"port"`
}

func LoadConfig(path string) (config Config, err error) {
	if path == "" {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

	} else {
		viper.SetConfigFile(path)
	}

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
