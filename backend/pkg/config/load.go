package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type loadConfig struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		Url      string `mapstructure:"url"`
	} `mapstructure:"database"`

	Logger struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"logger"`
}

func LoadConfig(cfg string) (*loadConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found")
		} else {
			return nil, err
		}
	}

	config := &loadConfig{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
