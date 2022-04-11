package config

import (
	"fmt"
	"strconv"
)

type ServerConfig struct {
	Name  string
	Port  int
	Https bool
}

func NewServerConfig() (ServerConfig, error) {
	var cfg ServerConfig
	serverPort, err := strconv.Atoi(GetEnv("SERVER_PORT", "8900"))
	if err != nil {
		return cfg, fmt.Errorf("unable to parse server port: %s", err.Error())
	}

	cfg = ServerConfig{
		Port:  serverPort,
		Name:  GetEnv("SERVER_NAME", "go-app"),
		Https: GetEnv("SERVER_HTTPS", "false") == "true",
	}

	return cfg, nil
}
