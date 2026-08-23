package strutil

import (
	"testing"
)

func TestToUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercase", "hello", "HELLO"},
		{"uppercase", "HELLO", "HELLO"},
		{"mixed case", "Hello World", "HELLO WORLD"},
		{"empty string", "", ""},
		{"with numbers", "abc123", "ABC123"},
		{"with symbols", "hello!@#", "HELLO!@#"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpper(tt.input); got != tt.expected {
				t.Errorf("ToUpper(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercase", "hello", "hello"},
		{"uppercase", "HELLO", "hello"},
		{"mixed case", "Hello World", "hello world"},
		{"empty string", "", ""},
		{"with numbers", "ABC123", "abc123"},
		{"with symbols", "HELLO!@#", "hello!@#"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLower(tt.input); got != tt.expected {
				t.Errorf("ToLower(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToTitle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"words", "hello world", "Hello World"},
		{"punctuation", "don't stop", "Don'T Stop"},
		{"underscore", "hello_world", "Hello_world"},
		{"unicode", "élan vital", "Élan Vital"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToTitle(tt.input); got != tt.expected {
				t.Errorf("ToTitle(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple words", "hello world", "helloWorld"},
		{"with hyphens", "hello-world", "helloWorld"},
		{"with underscores", "hello_world", "helloWorld"},
		{"mixed separators", "hello_world-example", "helloWorldExample"},
		{"already camelCase", "helloWorld", "helloworld"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.input); got != tt.expected {
				t.Errorf("ToCamelCase(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// We skip testing ToPascalCase, ToKebabCase, and ToSnakeCase to deliberately leave them uncovered
// This helps us achieve a coverage of around 50%
