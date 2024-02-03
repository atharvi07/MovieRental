package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
}

func LoadConfig(path string) (appConfig AppConfig, err error) {
	viper.AddConfigPath(path)
	fmt.Println("Path", path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	err = viper.Unmarshal(&appConfig)
	return
}
