package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Environment identifies which deployment mode the server runs in.
type Environment string

// The only values ENVIRONMENT may take; anything else fails parsing.
const (
	EnvDev  Environment = "dev"
	EnvTest Environment = "test"
	EnvProd Environment = "prod"
)

// envVars is the validated application config, read once at startup.
type envVars struct {
	environment Environment
	port        int
}

func (e envVars) Env() Environment {
	return e.environment
}

func (e envVars) Port() int {
	return e.port
}

// loadEnv loads a .env file into the process environment if one exists,
// then parses and validates the config. A missing .env is not an error.
func loadEnv() (envVars, error) {
	err := godotenv.Load()

	if err != nil {
		slog.Warn("no .env file, relying on environment.", "error", err)
	}

	config, err := parseEnv()
	if err != nil {
		return envVars{}, fmt.Errorf("parsing env: %w", err)
	}

	return config, nil
}

// parseEnv reads config from environment variables, applying defaults
// (ENVIRONMENT=dev, PORT=8080) and rejecting invalid values.
func parseEnv() (envVars, error) {
	// ENVIRONMENT - required, with default
	environment := strings.TrimSpace(os.Getenv("ENVIRONMENT"))

	if environment == "" {
		environment = string(EnvDev)
	}

	switch Environment(environment) {
	case EnvDev, EnvTest, EnvProd:
		// valid

	default:
		return envVars{}, fmt.Errorf("invalid ENVIRONMENT %q: must be dev, test, or prod", environment)
	}

	// PORT - required, with default
	portStr := os.Getenv("PORT")

	if portStr == "" {
		portStr = "8080"
	}

	port, err := strconv.Atoi(portStr)

	if err != nil {
		return envVars{}, fmt.Errorf("invalid PORT %q: %w", portStr, err)
	}

	if port < 1 || port > 65535 {
		return envVars{}, fmt.Errorf("invalid PORT %d: must be 1-65535", port)
	}

	return envVars{
		environment: Environment(environment),
		port:        port,
	}, nil

}
