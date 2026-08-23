package strutil

import (
	"strings"
	"unicode"
)

// ToUpper converts a string to uppercase
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower converts a string to lowercase
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToTitle converts a string to title case
func ToTitle(s string) string {
	previous := ' '
	return strings.Map(func(r rune) rune {
		if isTitleSeparator(previous) {
			r = unicode.ToTitle(r)
		}
		previous = r
		return r
	}, s)
}

func isTitleSeparator(r rune) bool {
	if r <= unicode.MaxASCII {
		return r != '_' && !unicode.IsLetter(r) && !unicode.IsDigit(r)
	}
	return unicode.IsSpace(r)
}

// ToCamelCase converts a string to camelCase
func ToCamelCase(s string) string {
	// Split the string by spaces, hyphens, and underscores
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})

	if len(words) == 0 {
		return ""
	}

	// First word starts with lowercase
	result := strings.ToLower(words[0])

	// Remaining words start with uppercase
	for i := 1; i < len(words); i++ {
		if len(words[i]) > 0 {
			result += string(unicode.ToUpper(rune(words[i][0]))) + strings.ToLower(words[i][1:])
		}
	}

	return result
}

// ToPascalCase converts a string to PascalCase
func ToPascalCase(s string) string {
	// Split the string by spaces, hyphens, and underscores
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})

	var result string

	// All words start with uppercase
	for _, word := range words {
		if len(word) > 0 {
			result += string(unicode.ToUpper(rune(word[0]))) + strings.ToLower(word[1:])
		}
	}

	return result
}

// ToKebabCase converts a string to kebab-case
func ToKebabCase(s string) string {
	// Split the string by spaces, hyphens, and underscores
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})

	// Join with hyphens
	return strings.ToLower(strings.Join(words, "-"))
}

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	// Split the string by spaces, hyphens, and underscores
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})

	// Join with underscores
	return strings.ToLower(strings.Join(words, "_"))
}
