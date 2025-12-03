package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort int
	DB       DB
}

type DB struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	v := viper.New()
	v.AutomaticEnv()

	return &Config{
		HTTPPort: v.GetInt("HTTP_PORT"),
		DB: DB{
			User:     v.GetString("POSTGRES_USER"),
			Password: v.GetString("POSTGRES_PASSWORD"),
			Host:     v.GetString("POSTGRES_HOST"),
			Port:     v.GetInt("POSTGRES_PORT"),
			DB:       v.GetString("POSTGRES_DB"),
		},
	}
}
