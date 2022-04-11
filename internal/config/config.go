package config

import "os"

type Config struct {
	Server     ServerConfig
	DB         DbConfig
	ApiLimiter ApiLimiterConfig
	Cors       CorsConfig
	Mail       MailConfig
}

func New() (Config, error) {
	var cfg Config

	serverCfg, err := NewServerConfig()
	if err != nil {
		return cfg, err
	}

	dbCfg, err := NewDbConfig()
	if err != nil {
		return cfg, err
	}

	limiterCfg, err := NewApiLimiterConfig()
	if err != nil {
		return cfg, err
	}

	corsCfg, err := NewCorsConfig()
	if err != nil {
		return cfg, err
	}

	mailCfg, err := NewMailConfig()
	if err != nil {
		return cfg, err
	}

	cfg = Config{
		Server:     serverCfg,
		DB:         dbCfg,
		ApiLimiter: limiterCfg,
		Cors:       corsCfg,
		Mail:       mailCfg,
	}

	return cfg, nil
}

func GetEnv(key, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}

	return value
}
