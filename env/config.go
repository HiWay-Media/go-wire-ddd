package env

import (
	"github.com/caarlos0/env/v6"
)

type Configuration struct {
	AppEnv              string `env:"APP_ENV"`
	AppName             string `env:"APP_NAME"`
	LogLevel            string `env:"LOG_LEVEL"`
	RunningMode         string `env:"RUNNING_MODE"` //fallback or main
	DbUsername          string `env:"DB_USERNAME"`
	DbPassword          string `env:"DB_PASSWORD"`
	DbHost              string `env:"DB_HOST"`
	DbPort              string `env:"DB_PORT"`
	DbIdleConn          string `env:"DB_IDLE_CONN"`
	DbMaxConn           string `env:"DB_MAX_CONN"`
	DbName              string `env:"DB_NAME"`
)