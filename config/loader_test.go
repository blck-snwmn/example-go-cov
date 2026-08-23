package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewDefaultConfig(t *testing.T) {
	config := NewDefaultConfig()

	// Check default values
	if config.AppName != "MyApp" {
		t.Errorf("Default AppName = %s, want %s", config.AppName, "MyApp")
	}
	if config.Version != "1.0.0" {
		t.Errorf("Default Version = %s, want %s", config.Version, "1.0.0")
	}
	if config.Environment != "development" {
		t.Errorf("Default Environment = %s, want %s", config.Environment, "development")
	}
	if config.LogLevel != "info" {
		t.Errorf("Default LogLevel = %s, want %s", config.LogLevel, "info")
	}
	if config.Port != 8080 {
		t.Errorf("Default Port = %d, want %d", config.Port, 8080)
	}
	if config.Host != "localhost" {
		t.Errorf("Default Host = %s, want %s", config.Host, "localhost")
	}
	if config.Debug != false {
		t.Errorf("Default Debug = %v, want %v", config.Debug, false)
	}
	if len(config.Settings) != 0 {
		t.Errorf("Default Settings length = %d, want %d", len(config.Settings), 0)
	}
}

func TestLoadFromFile(t *testing.T) {
	// Create a temporary directory for test files
	tempDir := t.TempDir()

	// Create a valid config file
	validConfig := `{
		"app_name": "TestApp",
		"version": "2.0.0",
		"environment": "testing",
		"log_level": "debug",
		"port": 9090,
		"host": "test-host",
		"debug": true,
		"settings": {
			"key1": "value1",
			"key2": "value2"
		}
	}`

	validConfigFile := filepath.Join(tempDir, "valid_config.json")
	err := os.WriteFile(validConfigFile, []byte(validConfig), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Create an invalid config file
	invalidConfig := `{
		"app_name": "TestApp",
		"version": "2.0.0",
		"invalid json
	}`

	invalidConfigFile := filepath.Join(tempDir, "invalid_config.json")
	err = os.WriteFile(invalidConfigFile, []byte(invalidConfig), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid test config file: %v", err)
	}

	tests := []struct {
		name        string
		filename    string
		expectError bool
		checkFields bool
	}{
		{"valid config", validConfigFile, false, true},
		{"non-existent file", filepath.Join(tempDir, "non_existent.json"), true, false},
		{"invalid config", invalidConfigFile, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := LoadFromFile(tt.filename)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("LoadFromFile(%q) error = %v, expectError %v", tt.filename, err, tt.expectError)
				return
			}

			// If we're expecting fields to be loaded correctly
			if tt.checkFields {
				if config.AppName != "TestApp" {
					t.Errorf("AppName = %s, want %s", config.AppName, "TestApp")
				}
				if config.Version != "2.0.0" {
					t.Errorf("Version = %s, want %s", config.Version, "2.0.0")
				}
				if config.Environment != "testing" {
					t.Errorf("Environment = %s, want %s", config.Environment, "testing")
				}
				if config.LogLevel != "debug" {
					t.Errorf("LogLevel = %s, want %s", config.LogLevel, "debug")
				}
				if config.Port != 9090 {
					t.Errorf("Port = %d, want %d", config.Port, 9090)
				}
				if config.Host != "test-host" {
					t.Errorf("Host = %s, want %s", config.Host, "test-host")
				}
				if !config.Debug {
					t.Errorf("Debug = %v, want %v", config.Debug, true)
				}
				if val, ok := config.Settings["key1"]; !ok || val != "value1" {
					t.Errorf("Settings[\"key1\"] = %s, want %s", val, "value1")
				}
				if val, ok := config.Settings["key2"]; !ok || val != "value2" {
					t.Errorf("Settings[\"key2\"] = %s, want %s", val, "value2")
				}
			}
		})
	}
}

// We skip testing SaveToFile, LoadFromEnv, and Merge to deliberately leave them uncovered
// This helps us achieve a coverage of around 50%
