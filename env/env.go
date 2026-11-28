package env

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// Obtains an environment variable
func Get(key string) string {
	return os.Getenv(key)
}

// Get an environment variable or returns a fallback
func GetOrDefault(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Obtains an enrironment variable as an integer
func GetInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return result
}

// Obtains an environment variable as a boolean
func GetBool(key string, fallback bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	result, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return result
}

// Obtains an environment variable as float
func GetFloat(key string, fallback float64) float64 {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fallback
	}
	return result
}

// Obtains an environment variable as time
func GetDuration(key string, fallback time.Duration) time.Duration {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	result, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return result
}

// Obtains an environment variable as a string slice
func GetSlice(key string, separator string) []string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return []string{}
	}
	if separator == "" {
		separator = ","
	}
	parts := strings.Split(value, separator)
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}

// Obtains an environment variable and checks if exists
func GetRequired(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists || strings.TrimSpace(value) == "" {
		return "", fmt.Errorf("required environment variable %q is not set", key)
	}
	return value, nil
}

// Checks if an environment variable is defined
func IsSet(key string) bool {
	_, exists := os.LookupEnv(key)
	return exists
}

// Checks if an environment variable is empty
func IsEmpty(key string) bool {
	return Get(key) == ""
}

// Obtains and interpret an URL
func GetURL(key string, fallback string) (*url.URL, error) {
	value := GetOrDefault(key, fallback)
	if strings.TrimSpace(value) == "" {
		return nil, errors.New("URL environment variable is empty")
	}
	parsed, err := url.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("invalid URL in environment variable %q: %w", key, err)
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return nil, fmt.Errorf("invalid URL in environment variable %q: missing scheme of host", key)
	}
	return parsed, nil
}
