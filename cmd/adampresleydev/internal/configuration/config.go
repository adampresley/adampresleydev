package configuration

import (
	"github.com/adampresley/adamgokit/mux2"
	"github.com/adampresley/configinator"
)

type Config struct {
	mux2.Config
	LogLevel string `flag:"loglevel" env:"LOG_LEVEL" default:"debug" description:"The log level to use. Valid values are 'debug', 'info', 'warn', and 'error'"`
}

func LoadConfig() Config {
	config := Config{}
	configinator.Behold(&config)
	return config
}
