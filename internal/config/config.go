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

type Config struct {
	Server   Server          `env:"SERVER"`
	DataBase database.Config `env:"POSTGRES"`
	Redis    redis.Config    `env:"REDIS"`
	Logger   logger.Config   `env:"LOGGER"`
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
	return &cfg, nil
}
