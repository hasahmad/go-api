package config

import (
	"fmt"
	"strconv"
)

// ApiLimiterConfig limit number of api requests for session
type ApiLimiterConfig struct {
	RPS     float64 // Rate limiter maximum requests per second
	Burst   int     // Rate limiter maximum burst
	Enabled bool    // Enable rate limiter
}

func NewApiLimiterConfig() (ApiLimiterConfig, error) {
	var cfg ApiLimiterConfig
	limiterRps, err := strconv.ParseFloat(GetEnv("LIMITER_RPS", "2"), 32)
	if err != nil {
		return cfg, fmt.Errorf("unable to parse limiter rps: %s", err.Error())
	}
	limiterBurst, err := strconv.Atoi(GetEnv("LIMITER_BURST", "4"))
	if err != nil {
		return cfg, fmt.Errorf("unable to parse limiter burst: %s", err.Error())
	}

	cfg = ApiLimiterConfig{
		RPS:     limiterRps,
		Burst:   limiterBurst,
		Enabled: GetEnv("LIMITER_ENABLED", "true") == "true",
	}

	return cfg, nil
}
