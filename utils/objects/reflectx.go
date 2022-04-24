package objects

import "reflect"

// RequireNonNil returns the value stored in the given optional,
func RequireNonNil[T any](value T) T {
	if IsNil(value) {
		panic("value is nil")
	}
	return value
}

func IsNil[T any](value T) bool {
	typeNil := reflect.TypeOf(value) == nil
	valueOf := reflect.ValueOf(value)
	return typeNil || !valueOf.IsValid() || (valueOf.Kind() == reflect.Ptr && valueOf.IsNil())
}
