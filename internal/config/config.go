package config

import (
	"os"
	"strconv"
)

// Config holds application configuration
type Config struct {
	ServerPort   string
	OpenAIKey    string
	AnthropicKey string
	Debug        bool
	Timeout      int
}

// NewConfig creates a new configuration from environment variables
func NewConfig() *Config {
	return &Config{
		ServerPort:   getEnv("PORT", "8080"),
		OpenAIKey:    getEnv("OPENAI_API_KEY", ""),
		AnthropicKey: getEnv("ANTHROPIC_API_KEY", ""),
		Debug:        getEnvAsBool("DEBUG", false),
		Timeout:      getEnvAsInt("TIMEOUT", 30),
	}
}

// getEnv returns the value of an environment variable or a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsBool returns the value of an environment variable as a bool or a default value
func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

// getEnvAsInt returns the value of an environment variable as an int or a default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
