package config

import (
	"fmt"
	"strconv"
)

type DbConfig struct {
	Type         string
	Host         string
	Port         int
	Name         string
	User         string
	Pass         string
	SSlMode      string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func NewDbConfig() (DbConfig, error) {
	var cfg DbConfig
	dbPort, err := strconv.Atoi(GetEnv("DB_PORT", "5432"))
	if err != nil {
		return cfg, fmt.Errorf("unable to parse db port: %s", err.Error())
	}
	dbMaxOpenConns, err := strconv.Atoi(GetEnv("DB_MAX_OPEN_CONNS", "25"))
	if err != nil {
		return cfg, fmt.Errorf("unable to parse db port: %s", err.Error())
	}
	dbMaxIdleConns, err := strconv.Atoi(GetEnv("DB_MAX_IDLE_CONNS", "25"))
	if err != nil {
		return cfg, fmt.Errorf("unable to parse db port: %s", err.Error())
	}
	cfg = DbConfig{
		Type:         GetEnv("DB_TYPE", "postgres"),
		Host:         GetEnv("DB_HOST", "localhost"),
		Port:         dbPort,
		Name:         GetEnv("DB_NAME", "dbname"),
		User:         GetEnv("DB_USER", "user"),
		Pass:         GetEnv("DB_PASS", "pass"),
		SSlMode:      GetEnv("DB_SSLMODE", "disable"),
		MaxOpenConns: dbMaxOpenConns,
		MaxIdleConns: dbMaxIdleConns,
		MaxIdleTime:  GetEnv("DB_MAX_IDLE_TIME", "15m"),
	}

	return cfg, nil
}
