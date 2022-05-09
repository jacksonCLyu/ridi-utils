package objects

import "reflect"

// RequireNonNil returns the value of the given optional struct, or panics if the value is not set.
func RequireNonNil[T any](value T) T {
	if IsNil(value) {
		panic("value is nil")
	}
	return value
}

// RequireNonNilM returns the value of the given optional struct, or panics if the value is not set with the given message.
func RequireNonNilM[T any](value T, message string) T {
	if IsNil(value) {
		panic(message)
	}
	return value
}

// IsNil returns true if the given optional struct is not set.
func IsNil[T any](value T) bool {
	vof := reflect.ValueOf(value)
	if vof.Kind() == reflect.Ptr {
		return vof.IsNil()
	}
	return false
}
