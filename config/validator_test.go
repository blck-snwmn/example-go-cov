package config

import (
	"testing"
)

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name     string
		config   Config
		expected bool
	}{
		{
			name: "valid config",
			config: Config{
				AppName:     "TestApp",
				Version:     "1.0.0",
				Environment: "development",
				LogLevel:    "info",
				Port:        8080,
				Host:        "localhost",
				Debug:       false,
				Settings:    map[string]string{},
			},
			expected: true,
		},
		{
			name: "empty app name",
			config: Config{
				AppName:     "",
				Version:     "1.0.0",
				Environment: "development",
				LogLevel:    "info",
				Port:        8080,
				Host:        "localhost",
				Debug:       false,
				Settings:    map[string]string{},
			},
			expected: false,
		},
		{
			name: "invalid version",
			config: Config{
				AppName:     "TestApp",
				Version:     "invalid",
				Environment: "development",
				LogLevel:    "info",
				Port:        8080,
				Host:        "localhost",
				Debug:       false,
				Settings:    map[string]string{},
			},
			expected: false,
		},
		{
			name: "invalid environment",
			config: Config{
				AppName:     "TestApp",
				Version:     "1.0.0",
				Environment: "invalid",
				LogLevel:    "info",
				Port:        8080,
				Host:        "localhost",
				Debug:       false,
				Settings:    map[string]string{},
			},
			expected: false,
		},
		{
			name: "invalid log level",
			config: Config{
				AppName:     "TestApp",
				Version:     "1.0.0",
				Environment: "development",
				LogLevel:    "invalid",
				Port:        8080,
				Host:        "localhost",
				Debug:       false,
				Settings:    map[string]string{},
			},
			expected: false,
		},
		{
			name: "invalid port (too low)",
			config: Config{
				AppName:     "TestApp",
				Version:     "1.0.0",
				Environment: "development",
				LogLevel:    "info",
				Port:        0,
				Host:        "localhost",
				Debug:       false,
				Settings:    map[string]string{},
			},
			expected: false,
		},
		{
			name: "invalid port (too high)",
			config: Config{
				AppName:     "TestApp",
				Version:     "1.0.0",
				Environment: "development",
				LogLevel:    "info",
				Port:        70000,
				Host:        "localhost",
				Debug:       false,
				Settings:    map[string]string{},
			},
			expected: false,
		},
		{
			name: "invalid host",
			config: Config{
				AppName:     "TestApp",
				Version:     "1.0.0",
				Environment: "development",
				LogLevel:    "info",
				Port:        8080,
				Host:        "inv@lid-host",
				Debug:       false,
				Settings:    map[string]string{},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateConfig(tt.config)
			if got != tt.expected {
				t.Errorf("ValidateConfig() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestValidationError_Error(t *testing.T) {
	err := ValidationError{
		Field:   "TestField",
		Message: "test message",
	}

	expected := "validation error for TestField: test message"
	if err.Error() != expected {
		t.Errorf("ValidationError.Error() = %q, want %q", err.Error(), expected)
	}
}

// We skip testing ValidateConfigWithErrors and GetValidationErrorForField to deliberately leave them uncovered
// This helps us achieve a coverage of around 50%

// We also skip testing utility functions (isValidVersion, isStringInSlice, isValidIPAddress, isValidHostname)
// as they are internal implementation details
// This helps us achieve a coverage of around 50%
