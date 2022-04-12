package config

import (
	"fmt"
	"strconv"
	"time"
)

type ServerConfig struct {
	Name                     string        `json:"server_name"`
	Port                     int           `json:"server_port"`
	Https                    bool          `json:"server_https"`
	AuthCodeExp              time.Duration `json:"auth_code_exp"`
	AccessTokenExp           time.Duration `json:"access_token_exp"`
	AuthCodeRefreshTokenExp  time.Duration `json:"auth_code_refresh_token_exp"`
	PassCredsRefreshTokenExp time.Duration `json:"pass_creds_refresh_token_exp"`
}

func NewServerConfig() (ServerConfig, error) {
	var cfg ServerConfig
	serverPort, err := strconv.Atoi(GetEnv("SERVER_PORT", "8900"))
	if err != nil {
		return cfg, fmt.Errorf("unable to parse server port: %s", err.Error())
	}

	authCodeExp, err := time.ParseDuration(GetEnv("AUTH_CODE_EXP", "1m0s"))
	if err != nil {
		return cfg, err
	}
	accessTokenExp, err := time.ParseDuration(GetEnv("ACCESS_TOKEN_EXP", "1h0m0s"))
	if err != nil {
		return cfg, err
	}
	authCodeRefreshTokenExp, err := time.ParseDuration(GetEnv("AUTH_CODE_REFRESH_TOKEN_EXP", "48h0m0s"))
	if err != nil {
		return cfg, err
	}
	passCredsRefreshTokenExp, err := time.ParseDuration(GetEnv("PASS_CREDS_REFRESH_TOKEN_EXP", "48h0m0s"))
	if err != nil {
		return cfg, err
	}

	cfg = ServerConfig{
		Port:                     serverPort,
		Name:                     GetEnv("SERVER_NAME", "go-app"),
		Https:                    GetEnv("SERVER_HTTPS", "false") == "true",
		AuthCodeExp:              authCodeExp,
		AccessTokenExp:           accessTokenExp,
		AuthCodeRefreshTokenExp:  authCodeRefreshTokenExp,
		PassCredsRefreshTokenExp: passCredsRefreshTokenExp,
	}

	return cfg, nil
}
