package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort string `mapstructure:"HTTP_PORT"`
	GrpcPort string `mapstructure:"GRPC_PORT"`
	SqlDSN   string `mapstructure:"POSTGRES_DSN"`
	RedisURL string `mapstructure:"REDIS_URL"`
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
