package config

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

// ValidationError represents a configuration validation error
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface
func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error for %s: %s", e.Field, e.Message)
}

// ValidateConfig validates the configuration and returns true if it's valid
func ValidateConfig(config Config) bool {
	return len(ValidateConfigWithErrors(config)) == 0
}

// ValidateConfigWithErrors validates the configuration and returns validation errors
func ValidateConfigWithErrors(config Config) []ValidationError {
	var errors []ValidationError

	// Validate AppName
	if strings.TrimSpace(config.AppName) == "" {
		errors = append(errors, ValidationError{
			Field:   "AppName",
			Message: "cannot be empty",
		})
	}

	// Validate Version (simple semver check)
	if !isValidVersion(config.Version) {
		errors = append(errors, ValidationError{
			Field:   "Version",
			Message: "not a valid version format (expected semver like x.y.z)",
		})
	}

	// Validate Environment
	validEnvs := []string{"development", "testing", "staging", "production"}
	if !isStringInSlice(strings.ToLower(config.Environment), validEnvs) {
		errors = append(errors, ValidationError{
			Field:   "Environment",
			Message: fmt.Sprintf("must be one of: %s", strings.Join(validEnvs, ", ")),
		})
	}

	// Validate LogLevel
	validLogLevels := []string{"debug", "info", "warn", "error", "fatal"}
	if !isStringInSlice(strings.ToLower(config.LogLevel), validLogLevels) {
		errors = append(errors, ValidationError{
			Field:   "LogLevel",
			Message: fmt.Sprintf("must be one of: %s", strings.Join(validLogLevels, ", ")),
		})
	}

	// Validate Port
	if config.Port < 1 || config.Port > 65535 {
		errors = append(errors, ValidationError{
			Field:   "Port",
			Message: "must be between 1 and 65535",
		})
	}

	// Validate Host
	if config.Host != "localhost" && !isValidIPAddress(config.Host) && !isValidHostname(config.Host) {
		errors = append(errors, ValidationError{
			Field:   "Host",
			Message: "not a valid hostname or IP address",
		})
	}

	return errors
}

// GetValidationErrorForField returns the validation error for a specific field, if any
func GetValidationErrorForField(errors []ValidationError, field string) *ValidationError {
	for _, err := range errors {
		if err.Field == field {
			return &err
		}
	}
	return nil
}

// isValidVersion checks if a string is a valid semantic version
func isValidVersion(version string) bool {
	pattern := `^(\d+)\.(\d+)\.(\d+)(-[0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*)?(\+[0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*)?$`
	match, _ := regexp.MatchString(pattern, version)
	return match
}

// isStringInSlice checks if a string is in a slice
func isStringInSlice(s string, slice []string) bool {
	for _, item := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// isValidIPAddress checks if a string is a valid IP address
func isValidIPAddress(s string) bool {
	return net.ParseIP(s) != nil
}

// isValidHostname checks if a string is a valid hostname
func isValidHostname(s string) bool {
	// Hostname validation regex
	pattern := `^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])(\.([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9]))*$`
	match, _ := regexp.MatchString(pattern, s)
	return match && len(s) <= 255
}
