package strutil

import (
	"regexp"
	"strconv"
	"strings"
)

// ParseWords splits a string into words and returns them as a slice
func ParseWords(s string) []string {
	// Split by whitespace
	return strings.Fields(s)
}

// ParseInt parses a string as an integer with proper error handling
func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ParseFloat parses a string as a float with proper error handling
func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// ParseBool parses a string as a boolean
// Accepts "1", "t", "T", "true", "TRUE", "True", "y", "yes", "Y", "YES" as true
// Accepts "0", "f", "F", "false", "FALSE", "False", "n", "no", "N", "NO" as false
func ParseBool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "1", "t", "true", "y", "yes":
		return true, nil
	case "0", "f", "false", "n", "no":
		return false, nil
	default:
		return false, strconv.ErrSyntax
	}
}

// ExtractEmails extracts email addresses from a string
func ExtractEmails(s string) []string {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	return emailRegex.FindAllString(s, -1)
}

// ExtractURLs extracts URLs from a string
func ExtractURLs(s string) []string {
	urlRegex := regexp.MustCompile(`https?://[^\s]+`)
	return urlRegex.FindAllString(s, -1)
}

// ExtractNumbers extracts numeric values from a string
func ExtractNumbers(s string) []string {
	numberRegex := regexp.MustCompile(`[-+]?\d*\.?\d+`)
	return numberRegex.FindAllString(s, -1)
}

// CountWords counts the number of words in a string
func CountWords(s string) int {
	return len(ParseWords(s))
}

// HasPrefix checks if the string starts with the specified prefix
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix checks if the string ends with the specified suffix
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}
