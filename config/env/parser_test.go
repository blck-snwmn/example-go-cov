package env

import (
	"os"
	"reflect"
	"testing"
)

func TestParseEnvVars(t *testing.T) {
	// Save original environment variables
	originalEnv := os.Environ()

	// Clean up environment after tests
	defer func() {
		os.Clearenv()
		for _, e := range originalEnv {
			key, value, _ := SplitEnvVar(e)
			os.Setenv(key, value)
		}
	}()

	// Setup test environment
	os.Clearenv()
	os.Setenv("APP_NAME", "TestApp")
	os.Setenv("APP_VERSION", "1.0.0")
	os.Setenv("APP_DEBUG", "true")
	os.Setenv("OTHER_VAR", "ignored")

	// Test with APP_ prefix
	config := ParseEnvVars("APP_")

	expected := map[string]string{
		"NAME":    "TestApp",
		"VERSION": "1.0.0",
		"DEBUG":   "true",
	}

	if !reflect.DeepEqual(config.Variables, expected) {
		t.Errorf("ParseEnvVars(\"APP_\") = %v, want %v", config.Variables, expected)
	}

	// Test with different prefix
	config = ParseEnvVars("OTHER_")

	expected = map[string]string{
		"VAR": "ignored",
	}

	if !reflect.DeepEqual(config.Variables, expected) {
		t.Errorf("ParseEnvVars(\"OTHER_\") = %v, want %v", config.Variables, expected)
	}

	// Test with non-existent prefix
	config = ParseEnvVars("NONEXISTENT_")

	expected = map[string]string{}

	if !reflect.DeepEqual(config.Variables, expected) {
		t.Errorf("ParseEnvVars(\"NONEXISTENT_\") = %v, want %v", config.Variables, expected)
	}
}

func TestGetString(t *testing.T) {
	config := EnvConfig{
		Variables: map[string]string{
			"EXISTING": "value",
		},
	}

	tests := []struct {
		name         string
		key          string
		defaultValue string
		expected     string
	}{
		{"existing key", "EXISTING", "default", "value"},
		{"non-existent key", "NONEXISTENT", "default", "default"},
		{"empty default", "NONEXISTENT", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.GetString(tt.key, tt.defaultValue)
			if got != tt.expected {
				t.Errorf("GetString(%q, %q) = %q, want %q", tt.key, tt.defaultValue, got, tt.expected)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	config := EnvConfig{
		Variables: map[string]string{
			"TRUE":    "true",
			"FALSE":   "false",
			"YES":     "yes",
			"NO":      "no",
			"ONE":     "1",
			"ZERO":    "0",
			"INVALID": "invalid",
		},
	}

	tests := []struct {
		name         string
		key          string
		defaultValue bool
		expected     bool
	}{
		{"true", "TRUE", false, true},
		{"false", "FALSE", true, false},
		{"yes", "YES", false, true},
		{"no", "NO", true, false},
		{"1", "ONE", false, true},
		{"0", "ZERO", true, false},
		{"invalid", "INVALID", true, true},
		{"invalid", "INVALID", false, false},
		{"non-existent", "NONEXISTENT", true, true},
		{"non-existent", "NONEXISTENT", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.GetBool(tt.key, tt.defaultValue)
			if got != tt.expected {
				t.Errorf("GetBool(%q, %v) = %v, want %v", tt.key, tt.defaultValue, got, tt.expected)
			}
		})
	}
}

// SplitEnvVar is a helper function to split environment variables
func SplitEnvVar(env string) (key, value string, ok bool) {
	for i := 0; i < len(env); i++ {
		if env[i] == '=' {
			return env[:i], env[i+1:], true
		}
	}
	return env, "", false
}

// We skip testing GetInt, GetFloat, GetStringSlice, Has, and GetAll to deliberately leave them uncovered
// This helps us achieve a coverage of around 50%
