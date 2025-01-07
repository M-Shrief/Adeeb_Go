package utils

import "reflect"

func IsString(value interface{}) bool {
	t := reflect.TypeOf(value)
	return t.Kind() == reflect.String
}

func IsBool(value interface{}) bool {
	t := reflect.TypeOf(value)
	return t.Kind() == reflect.Bool
}

func IsNumber(value interface{}) bool {
	t := reflect.TypeOf(value)
	return t.Kind() == reflect.Int || t.Kind() == reflect.Uint ||
		t.Kind() == reflect.Int8 || t.Kind() == reflect.Uint8 ||
		t.Kind() == reflect.Int16 || t.Kind() == reflect.Uint16 ||
		t.Kind() == reflect.Int32 || t.Kind() == reflect.Uint32 ||
		t.Kind() == reflect.Int64 || t.Kind() == reflect.Uint64 ||
		t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64
}
