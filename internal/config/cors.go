package config

import (
	"strings"
)

type CorsConfig struct {
	TrustedOrigins []string
}

func NewCorsConfig() (CorsConfig, error) {
	var cfg CorsConfig
	trustedOrigins := GetEnv("CORS_TRUSTED_ORIGINS", "localhost")

	cfg = CorsConfig{
		TrustedOrigins: strings.Split(trustedOrigins, ","),
	}

	return cfg, nil
}
