package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mohammadne/middleman/internal/network"
	"github.com/mohammadne/middleman/pkg/logger"
)

type client struct {
	Client *network.ClientConfig
	Logger *logger.Config
}

func Client(env string) *client {
	config := &client{}

	switch env {
	case "prod":
		config.loadProd()
	default:
		config.loadDev()
	}

	return config
}

func (config *client) loadProd() {
	config.Logger = &logger.Config{}

	// process
	envconfig.MustProcess("client", config)
	envconfig.MustProcess("client_client", config.Logger)
	envconfig.MustProcess("client_logger", config.Logger)
}

func (config *client) loadDev() {
	config.Client = &network.ClientConfig{
		RequestsNumber:   100,
		RequestsInterval: 200,
	}

	config.Logger = &logger.Config{
		Development:      true,
		EnableCaller:     true,
		EnableStacktrace: false,
		Encoding:         "console",
		Level:            "warn",
	}
}
