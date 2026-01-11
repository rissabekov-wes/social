package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
)

type EnvConfig struct {
	ServiceName string `env:"SERVICE_NAME"`
	// ServiceDomain string `env:"SERVICE_DOMAIN"`
	ServerPort int `env:"SERVER_PORT" envDefault:"8081"`
	// EnvName       string `env:"ENV_NAME" envDefault:"local"`

	// ReadTimeout       int `env:"READ_TIMEOUT" envDefault:"20"`
	// WriteTimeout      int `env:"WRITE_TIMEOUT" envDefault:"40"`
	// IdleTimeout       int `env:"IDLE_TIMEOUT" envDefault:"60"`
	// ReadHeaderTimeout int `env:"READ_HEADER_TIMEOUT" envDefault:"10"`

	DBMaxOpenConns int    `env:"DB_MAX_OPEN_CONNS" envDefault:"25"`
	DBMaxIdleConns int    `env:"DB_MAX_IDLE_CONNS" envDefault:"25"`
	DBMaxIdleTime  string `env:"DB_MAX_IDLE_TIME" envDefault:"15m"`

	DBAddr string `env:"DB_ADDR" envDefault:"postgres://user:pass@localhost:5432/social?sslmode=disable"`
}

func NewEnvironmentConfig() *EnvConfig {
	cfg := &EnvConfig{}
	if err := env.Parse(cfg); err != nil {
		panic(fmt.Sprintf("cannot find configs for server: %v", err))
	}

	log.Printf("Loaded environment config: %+v", cfg)

	return cfg
}
