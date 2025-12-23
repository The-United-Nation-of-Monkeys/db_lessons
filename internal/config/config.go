package config

import (
	"fmt"
	"os"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/database"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/redis"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Server struct {
	Version  int    `env:"SERVER_VERSION"`
	PortHttp uint16 `env:"SERVER_PORT_HTTP" default:"8080"`
}

type JWTConfig struct {
	PrivateKeyPath string `env:"JWT_PRIVATE_KEY_PATH" env-default:"keys/private.pem"`
	PublicKeyPath  string `env:"JWT_PUBLIC_KEY_PATH" env-default:"keys/public.pem"`
	RefreshTimeExp string `env:"JWT_REFRESH_TIME_EXP" env-default:"168h"`
	AccessTimeExp  string `env:"JWT_ACCESS_TIME_EXP" env-default:"1h"`
}

func (j *JWTConfig) GetPrivateKeyPath() string {
	if j.PrivateKeyPath == "" {
		return "keys/private.pem"
	}
	return j.PrivateKeyPath
}

func (j *JWTConfig) GetPublicKeyPath() string {
	if j.PublicKeyPath == "" {
		return "keys/public.pem"
	}
	return j.PublicKeyPath
}

type Config struct {
	Server   Server          `env:"SERVER"`
	DataBase database.Config `env:"POSTGRES"`
	Redis    redis.Config    `env:"REDIS"`
	Logger   logger.Config   `env:"LOGGER"`
	JWT      JWTConfig       `env-prefix:"JWT_"`
}

func New() (*Config, error) {
	var cfg Config

	if err := godotenv.Load("config.env"); err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("load config.env error: %w", err)
		}

		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("load .env error: %w", err)
		}
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	if cfg.JWT.PrivateKeyPath == "" {
		cfg.JWT.PrivateKeyPath = "keys/private.pem"
	}
	if cfg.JWT.PublicKeyPath == "" {
		cfg.JWT.PublicKeyPath = "keys/public.pem"
	}
	if cfg.JWT.RefreshTimeExp == "" {
		cfg.JWT.RefreshTimeExp = "168h"
	}
	if cfg.JWT.AccessTimeExp == "" {
		cfg.JWT.AccessTimeExp = "1h"
	}

	return &cfg, nil
}
