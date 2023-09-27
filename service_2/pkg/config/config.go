package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port           string `mapstructure:"PORT"`
	Service_1_port string `mapstructure:"SERVICE_1_PORT"`
}

func LoadConfig() (c *Config, err error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	return
}
