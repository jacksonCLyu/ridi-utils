package assignutil

import "github.com/jacksonCLyu/ridi-utils/utils/errcheck"

// Assign assigns the values.
func Assign[V any](v V, err error) V {
	errcheck.CheckAndPanic(err)
	return v
}