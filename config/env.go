package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type environment string

const (
	DEV  environment = "DEVELOPMENT"
	PROD environment = "PRODUCTION"
	// Add more as needed...
)

// List all possible environment variables here.
type EnvConfig struct {
	ENVIRONMENT environment
	PORT        string
}

// MustloadEnv will load env vars from a `.env` file.
// If a `.env `file is not provided it'll fallback to the runtime injected ones.
// It'll parse and validate the env and panic if fails to do so.
func MustloadEnv() *EnvConfig {
	// Discarding the error for later validation.
	godotenv.Load(".env")

	env := &EnvConfig{
		ENVIRONMENT: environment(os.Getenv("ENVIRONMENT")),
		PORT:        os.Getenv("PORT"),
	}
	parsedEnv, err := parseEnv(env)
	if err != nil {
		panic(err)
	}

	return parsedEnv
}

// parseEnv will validate the env vars received and return an error.
func parseEnv(env *EnvConfig) (*EnvConfig, error) {
	if env == nil {
		return nil, fmt.Errorf("No environment variables provided")
	}
	// Validate the possible environments.
	if env.ENVIRONMENT != DEV && env.ENVIRONMENT != PROD {
		return nil, fmt.Errorf("Invalid ENVIRONMENT %s", env.ENVIRONMENT)
	}
	if len(env.PORT) == 0 {
		// Defaulting to port 3000 if none provided.
		env.PORT = "3000"
	}

	return env, nil
}
