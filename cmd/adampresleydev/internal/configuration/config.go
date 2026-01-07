package configuration

import (
	"github.com/adampresley/configinator"
	"github.com/adampresley/mux"
)

type Config struct {
	mux.Config
	LogLevel string `flag:"loglevel" env:"LOG_LEVEL" default:"debug" description:"The log level to use. Valid values are 'debug', 'info', 'warn', and 'error'"`
}

func LoadConfig() Config {
	config := Config{}
	configinator.Behold(&config)
	return config
}
