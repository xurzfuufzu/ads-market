package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	Server Server `env-required:"true"`
	DB     DB     `env-required:"true"`
	JWT    JWT    `env-required:"true"`
}

type Server struct {
	Port string `env:"PORT" env-required:"true"`
}

type DB struct {
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     string `env:"DB_PORT" env-required:"true"`
	Username string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	Database string `env:"DB_NAME" env-required:"true"`
	URL      string
}

type JWT struct {
	SecretKey       string        `env:"JWT_SIGN_KEY" env-required:"true"`
	TokenTTL        time.Duration `env:"JWT_TOKEN_TTL" env-required:"true"`
	RefreshTokenTTL time.Duration `env:"JWT_REFRESH_TOKEN_TTL" env-required:"true"`
}

var instance *Config

func NewConfig() *Config {
	log.Println("Reading application config from environment...")

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	instance := &Config{}
	if err := cleanenv.ReadEnv(instance); err != nil {
		help, _ := cleanenv.GetDescription(instance, nil)
		log.Println(help)
		log.Fatal(err)
	}

	return instance
}
