package config

import "github.com/caarlos0/env/v6"

type Config struct {
	DB struct {
		Driver   string `env:"DB_DRIVER" envDefault:"postgres"`
		Host     string `env:"DB_HOST" envDefault:"localhost"`
		Port     int    `env:"DB_PORT" envDefault:"54320"`
		Name     string `env:"DB_NAME" envDefault:"portfolio"`
		User     string `env:"DB_USER" envDefault:"root"`
		Password string `env:"DB_PASSWORD" envDefault:"root"`
	}
	Server struct {
		Port int `env:"SERVER_PORT" envDefault:"8080"`
	}
}

func New() (*Config, error) {
	config := &Config{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}
