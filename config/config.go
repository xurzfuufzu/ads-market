package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type Config struct {
	Server Server `env-required:"true" yaml:"server"`
	DB     DB     `env-required:"true" yaml:"db"`
	JWT    JWT    `env-required:"true" yaml:"jwt"`
}

type Server struct {
	Port string `env-required:"true" yaml:"port"`
}

type DB struct {
	Host     string `env-required:"true" yaml:"host"`
	Port     string `env-required:"true" yaml:"port"`
	Username string `env-required:"true" yaml:"username"`
	Password string `env-required:"true" yaml:"password"`
	Database string `env-required:"true" yaml:"database"`
	URL      string `yaml:"-"`
}

type JWT struct {
	SecretKey       string        `env-required:"true" yaml:"sign_key"`
	TokenTTL        time.Duration `env-required:"true" yaml:"token_ttl"`
	RefreshTokenTTL time.Duration `env-required:"true" yaml:"refresh_token_ttl"`
}

var instance *Config

func NewConfig() *Config {
	log.Println("Reading application config...")

	instance := &Config{}
	if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
		help, _ := cleanenv.GetDescription(instance, nil)
		log.Println(help)
		log.Fatal(err)
	}

	return instance
}
