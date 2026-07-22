package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPPORT string `env:"HTTP_PORT" envDefault:"8080"`
	Services Services
}

type Services struct {
	UserService    UserService
	MovieService   MovieService
	BookingService BookingService
}

type UserService struct{
	Host string
	Port int
}

type MovieService struct{
	Host string
	Port int
}

type BookingService struct{
	Host string
	Port int
}

func New(path string)(*Config, error){
	var conf Config

	err := cleanenv.ReadConfig(path, &conf)
	if err != nil{
		return nil, fmt.Errorf("cleanenv.ReadConfig: %w", err)
	}

	return &conf, nil
}

