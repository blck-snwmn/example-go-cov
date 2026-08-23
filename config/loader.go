package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config represents a configuration object
type Config struct {
	// Common configuration fields
	AppName     string            `json:"app_name"`
	Version     string            `json:"version"`
	Environment string            `json:"environment"`
	LogLevel    string            `json:"log_level"`
	Port        int               `json:"port"`
	Host        string            `json:"host"`
	Debug       bool              `json:"debug"`
	Settings    map[string]string `json:"settings"`
}

// NewDefaultConfig creates a new Config with default values
func NewDefaultConfig() Config {
	return Config{
		AppName:     "MyApp",
		Version:     "1.0.0",
		Environment: "development",
		LogLevel:    "info",
		Port:        8080,
		Host:        "localhost",
		Debug:       false,
		Settings:    make(map[string]string),
	}
}

// LoadFromFile loads configuration from a JSON file
func LoadFromFile(filename string) (Config, error) {
	config := NewDefaultConfig()

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return config, fmt.Errorf("config file not found: %s", filename)
	}

	// Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	// Parse JSON
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}

// SaveToFile saves configuration to a JSON file
func SaveToFile(config Config, filename string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// Marshal config to JSON
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling config: %w", err)
	}

	// Write to file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

// LoadFromEnv loads configuration from environment variables
func LoadFromEnv(prefix string) Config {
	config := NewDefaultConfig()

	// Get all environment variables
	for _, env := range os.Environ() {
		// Check if environment variable starts with prefix
		if strings.HasPrefix(env, prefix) {
			// Split key and value
			parts := strings.SplitN(env, "=", 2)
			if len(parts) != 2 {
				continue
			}

			key := strings.TrimPrefix(parts[0], prefix)
			value := parts[1]

			// Set config value based on key
			switch strings.ToLower(key) {
			case "app_name", "appname":
				config.AppName = value
			case "version":
				config.Version = value
			case "environment", "env":
				config.Environment = value
			case "log_level", "loglevel":
				config.LogLevel = value
			case "host":
				config.Host = value
			case "port":
				if _, err := fmt.Sscanf(value, "%d", &config.Port); err != nil {
					continue
				}
			case "debug":
				config.Debug = strings.ToLower(value) == "true" || value == "1"
			default:
				// Store in settings
				config.Settings[key] = value
			}
		}
	}

	return config
}

// Merge merges two configurations, with the second taking precedence
func Merge(base, override Config) Config {
	result := base

	// Override fields if they are not empty
	if override.AppName != "" {
		result.AppName = override.AppName
	}
	if override.Version != "" {
		result.Version = override.Version
	}
	if override.Environment != "" {
		result.Environment = override.Environment
	}
	if override.LogLevel != "" {
		result.LogLevel = override.LogLevel
	}
	if override.Host != "" {
		result.Host = override.Host
	}
	if override.Port != 0 {
		result.Port = override.Port
	}

	// Debug is boolean, so it's always considered
	result.Debug = override.Debug

	// Merge settings
	for k, v := range override.Settings {
		result.Settings[k] = v
	}

	return result
}
