package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")
	configName := os.Getenv("CONFIG_NAME")
	if configPath == "" || configName == "" {
		log.Fatal("CONFIG_PATH or CONFIG_NAME environment variable not set")
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error unmarshalling config")
	}

	return &config
}
