package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
}

func (c AppConfig) config() *viper.Viper {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/goprom")
	viper.AddConfigPath("$HOME/.goprom")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file not found")
		} else {
			panic("Read config file error")
		}
	}
	return viper.GetViper()
}

func (c AppConfig) ApiBaseUrl() string {
	apiUrl := c.config().GetString("goprom.api.url")
	return apiUrl
}
func (c AppConfig) ApiGroup() string {
	apiGroup := c.config().GetString("goprom.api.group")
	return apiGroup
}

func NewConfig() AppConfig {
	return AppConfig{}
}
