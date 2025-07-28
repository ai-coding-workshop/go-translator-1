package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

// Config holds application configuration
type Config struct {
	ServerPort   string `yaml:"port"`
	OpenAIKey    string `yaml:"openai_key"`
	AnthropicKey string `yaml:"anthropic_key"`
	Debug        bool   `yaml:"debug"`
	Timeout      int    `yaml:"timeout"`
}

// NewConfig creates a new configuration from environment variables and config file
func NewConfig() (*Config, error) {
	// Parse command line flags
	configFile := flag.String("config", "", "Path to config file")
	flag.Parse()

	// Create default config
	config := &Config{
		ServerPort:   "8080",
		OpenAIKey:    "",
		AnthropicKey: "",
		Debug:        false,
		Timeout:      30,
	}

	// Load from config file if specified
	if *configFile != "" {
		if err := config.loadFromFile(*configFile); err != nil {
			return nil, fmt.Errorf("failed to load config from file: %w", err)
		}
	}

	// Override with environment variables
	config.loadFromEnv()

	return config, nil
}

// loadFromFile loads configuration from a YAML file
func (c *Config) loadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var fileConfig struct {
		Server struct {
			Port string `yaml:"port"`
		} `yaml:"server"`
		LLM struct {
			OpenAIKey    string `yaml:"openai_key"`
			AnthropicKey string `yaml:"anthropic_key"`
			Timeout      int    `yaml:"timeout"`
		} `yaml:"llm"`
		Debug bool `yaml:"debug"`
	}

	if err := yaml.Unmarshal(data, &fileConfig); err != nil {
		return err
	}

	// Apply file config
	if fileConfig.Server.Port != "" {
		c.ServerPort = fileConfig.Server.Port
	}
	if fileConfig.LLM.OpenAIKey != "" {
		c.OpenAIKey = fileConfig.LLM.OpenAIKey
	}
	if fileConfig.LLM.AnthropicKey != "" {
		c.AnthropicKey = fileConfig.LLM.AnthropicKey
	}
	if fileConfig.LLM.Timeout > 0 {
		c.Timeout = fileConfig.LLM.Timeout
	}
	c.Debug = fileConfig.Debug

	return nil
}

// loadFromEnv loads configuration from environment variables
func (c *Config) loadFromEnv() {
	if value := os.Getenv("PORT"); value != "" {
		c.ServerPort = value
	}
	if value := os.Getenv("OPENAI_API_KEY"); value != "" {
		c.OpenAIKey = value
	}
	if value := os.Getenv("ANTHROPIC_API_KEY"); value != "" {
		c.AnthropicKey = value
	}
	if value := os.Getenv("DEBUG"); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			c.Debug = boolValue
		}
	}
	if value := os.Getenv("TIMEOUT"); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			c.Timeout = intValue
		}
	}
}
