package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST string `mapstructure:"DB_HOST"`
	DB_PORT string `mapstructure:"DB_PORT"`
	DB_USER string `mapstructure:"DB_USER"`
	DB_NAME string `mapstructure:"DB_NAME"`
}

var (
	AppConfig *Config
)

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
