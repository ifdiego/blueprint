package validator

import (
	"fmt"
	"net/mail"
	"net/url"
	"regexp"
	"slices"
	"strings"
)

// Checks if a string is not empty
func Required(value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

// Validates and email address
func Email(value string) error {
	if err := Required(value); err != nil {
		return fmt.Errorf("invalid email: %w", err)
	}
	address, err := mail.ParseAddress(value)
	if err != nil || address.Address != value {
		return fmt.Errorf("invalid email address: %q", value)
	}
	return nil
}

// Validates a URL
func URL(value string) error {
	if err := Required(value); err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	parsed, err := url.ParseRequestURI(value)
	if err != nil {
		return fmt.Errorf("invalid URL: %q", value)
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return fmt.Errorf("URL must contain a scheme and host: %q", value)
	}
	return nil
}

// Validates the minimum string length
func MinLength(value string, min int) error {
	if min < 0 {
		return fmt.Errorf("minimum length cannot be negative")
	}
	if len([]rune(value)) < min {
		return fmt.Errorf("value must have at least %d characters", min)
	}
	return nil
}

// Validates the maximum string length
func MaxLength(value string, max int) error {
	if max < 0 {
		return fmt.Errorf("maximum length cannot be negative")
	}
	if len([]rune(value)) > max {
		return fmt.Errorf("value must have at most %d characters", max)
	}
	return nil
}

// Validates the string length range
func Length(value string, min, max int) error {
	if min < 0 || max < min {
		return fmt.Errorf("invalid length range: min=%d, max=%d", min, max)
	}
	size := len([]rune(value))
	if size < min || size > max {
		return fmt.Errorf("value must have between %d ad %d characters", min, max)
	}
	return nil
}

// Validates whether a value belongs to the allowed values
func In[T comparable](value T, allowed ...T) error {
	if slices.Contains(allowed, value) {
		return nil
	}
	return fmt.Errorf("value %v is not allowed", value)
}

// Validates a string against a regular expression
func Matches(value, pattern string) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("invalid regular expression: %w", err)
	}
	if !re.MatchString(value) {
		return fmt.Errorf("value %q does not match pattern %q", value, pattern)
	}
	return nil
}

// Validates a UUID string
func UUID(value string) error {
	pattern := `(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[1-8][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
	if err := Matches(value, pattern); err != nil {
		return fmt.Errorf("invalid UUID: %w", err)
	}
	return nil
}

// Validates whether a value belongs to a collection
func OneOf[T comparable](value T, allowed []T) error {
	return In(value, allowed...)
}
