package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	NETWORK string `env:"NETWORK"`
	ADDRESS string `env:"ADDRESS"`

	HTTPPort   string `env:"HTTP_PORT"`
	Storage    string `env:"STORAGE"`
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
}

func New(configPath string) (Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("cleanenv.ReadConfig: %w", err)
	}
	return cfg, nil //вместо err возвращаю nil, если всё ок
}