package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	GeminiAPIKey string `mapstructure:"GEMINI_API_KEY"`
}

func LoadConfig() (config Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Note: .env file not found, using system environment variables")
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Gagal unmarshal config: ", err)
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	return
}
