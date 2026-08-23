package json

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

// JSONObject represents a generic JSON object
type JSONObject map[string]interface{}

// JSONArray represents a generic JSON array
type JSONArray []interface{}

// Parse parses a JSON string into a JSONObject
func Parse(data string) (JSONObject, error) {
	var result JSONObject
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParseArray parses a JSON string into a JSONArray
func ParseArray(data string) (JSONArray, error) {
	var result JSONArray
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReadJSONFile reads a JSON file and parses it into a JSONObject
func ReadJSONFile(filename string) (result JSONObject, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := file.Close(); err == nil {
			err = closeErr
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return Parse(string(data))
}

// ReadJSONArrayFile reads a JSON file and parses it into a JSONArray
func ReadJSONArrayFile(filename string) (result JSONArray, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := file.Close(); err == nil {
			err = closeErr
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return ParseArray(string(data))
}

// WriteJSONFile writes a JSONObject to a file with proper indentation
func WriteJSONFile(filename string, data JSONObject, indent bool) error {
	var bytes []byte
	var err error

	if indent {
		bytes, err = json.MarshalIndent(data, "", "  ")
	} else {
		bytes, err = json.Marshal(data)
	}

	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0644)
}

// WriteJSONArrayFile writes a JSONArray to a file with proper indentation
func WriteJSONArrayFile(filename string, data JSONArray, indent bool) error {
	var bytes []byte
	var err error

	if indent {
		bytes, err = json.MarshalIndent(data, "", "  ")
	} else {
		bytes, err = json.Marshal(data)
	}

	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0644)
}

// GetString gets a string value from a JSONObject by key
func GetString(obj JSONObject, key string) (string, error) {
	val, ok := obj[key]
	if !ok {
		return "", errors.New("key not found")
	}

	str, ok := val.(string)
	if !ok {
		return "", errors.New("value is not a string")
	}

	return str, nil
}

// GetNumber gets a number value from a JSONObject by key
func GetNumber(obj JSONObject, key string) (float64, error) {
	val, ok := obj[key]
	if !ok {
		return 0, errors.New("key not found")
	}

	// JSON numbers are parsed as float64
	num, ok := val.(float64)
	if !ok {
		return 0, errors.New("value is not a number")
	}

	return num, nil
}

// GetBool gets a boolean value from a JSONObject by key
func GetBool(obj JSONObject, key string) (bool, error) {
	val, ok := obj[key]
	if !ok {
		return false, errors.New("key not found")
	}

	b, ok := val.(bool)
	if !ok {
		return false, errors.New("value is not a boolean")
	}

	return b, nil
}
