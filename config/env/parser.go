package env

import (
	"os"
	"strconv"
	"strings"
)

// EnvConfig represents a configuration parsed from environment variables
type EnvConfig struct {
	Variables map[string]string
}

// ParseEnvVars parses environment variables with the given prefix
func ParseEnvVars(prefix string) EnvConfig {
	config := EnvConfig{
		Variables: make(map[string]string),
	}

	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		if strings.HasPrefix(key, prefix) {
			// Remove prefix and store in variables map
			normalizedKey := strings.TrimPrefix(key, prefix)
			config.Variables[normalizedKey] = value
		}
	}

	return config
}

// GetString gets a string value from environment variables
func (c EnvConfig) GetString(key string, defaultValue string) string {
	value, exists := c.Variables[key]
	if !exists {
		return defaultValue
	}
	return value
}

// GetInt gets an integer value from environment variables
func (c EnvConfig) GetInt(key string, defaultValue int) int {
	value, exists := c.Variables[key]
	if !exists {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intValue
}

// GetBool gets a boolean value from environment variables
func (c EnvConfig) GetBool(key string, defaultValue bool) bool {
	value, exists := c.Variables[key]
	if !exists {
		return defaultValue
	}

	// Check for truthy values
	switch strings.ToLower(value) {
	case "true", "1", "yes", "y", "on":
		return true
	case "false", "0", "no", "n", "off":
		return false
	default:
		return defaultValue
	}
}

// GetFloat gets a float value from environment variables
func (c EnvConfig) GetFloat(key string, defaultValue float64) float64 {
	value, exists := c.Variables[key]
	if !exists {
		return defaultValue
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}

	return floatValue
}

// GetStringSlice gets a slice of strings from a comma-separated environment variable
func (c EnvConfig) GetStringSlice(key string, defaultValue []string) []string {
	value, exists := c.Variables[key]
	if !exists || value == "" {
		return defaultValue
	}

	return strings.Split(value, ",")
}

// Has checks if an environment variable exists
func (c EnvConfig) Has(key string) bool {
	_, exists := c.Variables[key]
	return exists
}

// GetAll returns all environment variables as a map
func (c EnvConfig) GetAll() map[string]string {
	// Return a copy to prevent modification of the internal map
	result := make(map[string]string)
	for k, v := range c.Variables {
		result[k] = v
	}
	return result
}
