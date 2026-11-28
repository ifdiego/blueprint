package assert

import (
	"reflect"
	"strings"
	"testing"
)

// Check if a condition is true
func True(t *testing.T, value bool) {
	t.Helper()
	if !value {
		t.Error("expected true, got false")
	}
}

// Check if a condition is false
func False(t *testing.T, value bool) {
	t.Helper()
	if value {
		t.Error("expected false, got true")
	}
}

// Check if a value is nil
func Nil(t *testing.T, value any) {
	t.Helper()
	if value == nil {
		return
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface:
		if v.IsNil() {
			return
		}
	}
	t.Errorf("expected nil, got: %v", value)
}

// Check if a value is not nil
func NotNil(t *testing.T, value any) {
	t.Helper()
	if value == nil {
		t.Errorf("expected non-nil, got nil")
		return
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface:
		if v.IsNil() {
			t.Error("expected non-nil, got nil")
		}
	}
}

// Compare two values
func Equal[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

// Check if two values are different
func NotEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected == actual {
		t.Errorf("expected values to be different, both are: %v", actual)
	}
}

// Checj if a value is empty
func Empty(t *testing.T, value any) {
	t.Helper()
	if reflect.ValueOf(value).IsValid() {
		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
			if v.Len() == 0 {
				return
			}
		}
	}
	t.Errorf("expected empty value, got: %v", value)
}

// Check if a value is not empty
func NotEmpty(t *testing.T, value any) {
	t.Helper()
	v := reflect.ValueOf(value)
	if !v.IsValid() {
		t.Error("expected non-empty value, got nil")
		return
	}
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		if v.Len() == 0 {
			t.Error("expected non-empty value")
		}
	default:
		t.Errorf("unsupported type for NotEmpty: %T", value)
	}
}

// Check if a slice contains a value
func Contains(t *testing.T, value, substring string) {
	t.Helper()
	if !strings.Contains(value, substring) {
		t.Errorf("expected %q to contain %q", value, substring)
	}
}

// Check if a function panics
func Panics(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, but function did not panic")
		}
	}()
	fn()
}

// Check if a function don't panic
func NoPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("expected no panic, but function panicked: %v", r)
		}
	}()
	fn()
}

// Check if a value is its type's zero value
func Zero[T comparable](t *testing.T, value T) {
	t.Helper()
	var zero T
	if value != zero {
		t.Errorf("expected zero value (%v), got: %v", zero, value)
	}
}

// Check if a value is not its type's zero value
func NotZero[T comparable](t *testing.T, value T) {
	t.Helper()
	var zero T
	if value == zero {
		t.Errorf("expectd non-zero value, got: %v", value)
	}
}
