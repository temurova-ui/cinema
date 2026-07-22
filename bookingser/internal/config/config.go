package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	NETWORK string `env:"NETWORK"`
	ADDRESS string `env:"ADDRESS"`

	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBBooking  string `env:"DB_BOOKING"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
}

func New(filepath string) (Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(filepath, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("cleanenv.ReadConfig %w", err)
	}
	return cfg, nil
}