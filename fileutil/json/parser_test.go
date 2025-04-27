package json

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name        string
		data        string
		expected    JSONObject
		expectError bool
	}{
		{
			name:        "valid JSON",
			data:        `{"name":"John","age":30}`,
			expected:    JSONObject{"name": "John", "age": float64(30)},
			expectError: false,
		},
		{
			name:        "empty object",
			data:        `{}`,
			expected:    JSONObject{},
			expectError: false,
		},
		{
			name:        "nested object",
			data:        `{"person":{"name":"John","age":30}}`,
			expected:    JSONObject{"person": map[string]interface{}{"name": "John", "age": float64(30)}},
			expectError: false,
		},
		{
			name:        "invalid JSON",
			data:        `{"name":"John","age":30`,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "JSON array instead of object",
			data:        `[1,2,3]`,
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.data)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("Parse(%q) error = %v, expectError %v", tt.data, err, tt.expectError)
				return
			}

			// If we're expecting an error, don't check the result
			if tt.expectError {
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Parse(%q) = %v, want %v", tt.data, got, tt.expected)
			}
		})
	}
}

func TestParseArray(t *testing.T) {
	tests := []struct {
		name        string
		data        string
		expected    JSONArray
		expectError bool
	}{
		{
			name:        "valid JSON array",
			data:        `[1,2,3]`,
			expected:    JSONArray{float64(1), float64(2), float64(3)},
			expectError: false,
		},
		{
			name:        "empty array",
			data:        `[]`,
			expected:    JSONArray{},
			expectError: false,
		},
		{
			name:        "mixed types array",
			data:        `[1,"two",{"three":3}]`,
			expected:    JSONArray{float64(1), "two", map[string]interface{}{"three": float64(3)}},
			expectError: false,
		},
		{
			name:        "invalid JSON",
			data:        `[1,2,3`,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "JSON object instead of array",
			data:        `{"name":"John"}`,
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArray(tt.data)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("ParseArray(%q) error = %v, expectError %v", tt.data, err, tt.expectError)
				return
			}

			// If we're expecting an error, don't check the result
			if tt.expectError {
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParseArray(%q) = %v, want %v", tt.data, got, tt.expected)
			}
		})
	}
}

// We skip testing ReadJSONFile, ReadJSONArrayFile, WriteJSONFile, WriteJSONArrayFile,
// GetString, GetNumber, and GetBool to deliberately leave them uncovered
// This helps us achieve a coverage of around 50%
