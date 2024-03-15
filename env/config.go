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
}



func GetEnvConfig() *Configuration {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		return
	}
	err := godotenv.Load(fmt.Sprintf("./env/.env.%s", strings.ToLower(appEnv)))
	if err == nil {
		return
	}
	_ = godotenv.Load()
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
    return &cfg
}