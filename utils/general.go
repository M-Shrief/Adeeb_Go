package utils

import "reflect"

func IsZeroValue[T any](x T) bool {
	zero := reflect.Zero(reflect.TypeOf(x)).Interface()
	return reflect.DeepEqual(x, zero)
}
