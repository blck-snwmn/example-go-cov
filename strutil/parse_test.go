package strutil

import (
	"reflect"
	"testing"
)

func TestParseWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"simple case", "hello world golang", []string{"hello", "world", "golang"}},
		{"extra spaces", "  hello  world  ", []string{"hello", "world"}},
		{"tabs and newlines", "hello\tworld\n golang", []string{"hello", "world", "golang"}},
		{"empty string", "", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseWords(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParseWords(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    int
		expectError bool
	}{
		{"positive number", "123", 123, false},
		{"negative number", "-123", -123, false},
		{"zero", "0", 0, false},
		{"invalid characters", "123abc", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInt(tt.input)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("ParseInt(%q) error = %v, expectError %v", tt.input, err, tt.expectError)
				return
			}

			// If we're expecting an error, don't check the result
			if tt.expectError {
				return
			}

			if got != tt.expected {
				t.Errorf("ParseInt(%q) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}

func TestParseBool(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    bool
		expectError bool
	}{
		{"true", "true", true, false},
		{"false", "false", false, false},
		{"TRUE", "TRUE", true, false},
		{"FALSE", "FALSE", false, false},
		{"t", "t", true, false},
		{"f", "f", false, false},
		{"1", "1", true, false},
		{"0", "0", false, false},
		{"yes", "yes", true, false},
		{"no", "no", false, false},
		{"invalid", "invalid", false, true},
		{"empty string", "", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBool(tt.input)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("ParseBool(%q) error = %v, expectError %v", tt.input, err, tt.expectError)
				return
			}

			// If we're expecting an error, don't check the result
			if tt.expectError {
				return
			}

			if got != tt.expected {
				t.Errorf("ParseBool(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

// We skip testing ParseFloat, ExtractEmails, ExtractURLs, ExtractNumbers, CountWords,
// HasPrefix, and HasSuffix to deliberately leave them uncovered
// This helps us achieve a coverage of around 50%
