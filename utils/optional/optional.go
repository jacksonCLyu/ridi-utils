package optional

import "github.com/jacksonCLyu/ridi-utils/utils/objects"

// Optional is a container for a value that may or may not be present.
type Optional[T any] struct {
	value T
}

func Of[T any](value T) *Optional[T] {
	return objects.RequireNonNil(&Optional[T]{value: value})
}

func OfNilable[T any](value T) *Optional[T] {
	if objects.IsNil(value) {
		return &Optional[T]{}
	} else {
		return Of(value)
	}
}

func (o *Optional[T]) Get() T {
	if objects.IsNil(o.value) {
		panic("no value present")
	}
	return o.value
}

func (o *Optional[T]) IsPresent() bool {
	return !objects.IsNil(o.value)
}

func (o *Optional[T]) IfPresent(consumer func(value T)) {
	if !objects.IsNil(o.value) {
		consumer(o.value)
	}
}

func (o *Optional[T]) filter(predict func(value T) bool) *Optional[T] {
	objects.RequireNonNil(predict)
	if !o.IsPresent() {
		return o
	} else {
		if predict(o.value) {
			return o
		} else {
			return &Optional[T]{}
		}
	}
}

// func (o *Optional[T]) Map[U any](mapper func(T) U) *Optional[U] {
// 	objects.RequireNonNil(mapper)
// 	if !o.IsPresent() {
// 		return &Optional[U]{}
// 	} else {
// 		return OfNilable(mapper(o.value))
// 	}
// }

func (o *Optional[T]) OrElse(other T) T {
	if !o.IsPresent() {
		return other
	} else {
		return o.value
	}
}

func (o *Optional[T]) OrElseGet(other func() T) T {
	if !o.IsPresent() {
		return other()
	} else {
		return o.value
	}
}

func (o *Optional[T]) OrElseThrow(exception func() error) T {
	if !o.IsPresent() {
		panic(exception())
	} else {
		return o.value
	}
}
