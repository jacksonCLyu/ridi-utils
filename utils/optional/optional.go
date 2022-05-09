package optional

import (
	"fmt"

	"github.com/jacksonCLyu/ridi-utils/utils/objects"
)

// optstruct is a struct that contains a value and a flag indicating whether the value is set.
type optstruct[T any] struct {
	value T
}

// Of returns an optional struct with the given value.
func Of[T any](value T) *optstruct[T] {
	return &optstruct[T]{value: objects.RequireNonNil(value)}
}

// OfNilable returns an optional struct with the given value.
func OfNilable[T any](value T) *optstruct[T] {
	return &optstruct[T]{value: value}
}

// Get returns the value of the optional struct.
func (o *optstruct[T]) Get() T {
	return o.value
}

// IsPresent returns true if the value is set.
func (o *optstruct[T]) IsPresent() bool {
	return !objects.IsNil(o.value)
}

// IfPresent executes the given function if the value is set.
func (o *optstruct[T]) IfPresent(consumer func(T)) {
	if o.IsPresent() {
		consumer(o.value)
	}
}

// Filter returns a new optional struct containing the value if the given predicate is true.
func (o *optstruct[T]) Filter(predicate func(T) bool) *optstruct[T] {
	objects.RequireNonNil(predicate)
	if o.IsPresent() {
		if predicate(o.value) {
			return o
		} else {
			return &optstruct[T]{}
		}
	} else {
		return o
	}
}

// OrElse returns the value of the given optional struct, or the given value if the value is not set.
func (o *optstruct[T]) OrElse(other T) T {
	if o.IsPresent() {
		return o.value
	} else {
		return other
	}
}

// OrElseGet returns the value of this optional struct, or the given value if this optional struct is not set.
func (o *optstruct[T]) OrElseGet(supplier func() T) T {
	if o.IsPresent() {
		return o.value
	} else {
		return supplier()
	}
}

// OrElsePanic throws an exception if the optional value is not set.
func (o *optstruct[T]) OrElsePanic(expSupplier func() any) T {
	if o.IsPresent() {
		return o.value
	} else {
		panic(expSupplier())
	}
}

// String returns the string representation of the optional value.
func (o *optstruct[T]) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("Optional[%v]", o.value)
	} else {
		return "Optional.empty"
	}
}
