package utils

import "reflect"

func GetDefaultValue(ptr interface{}, defaultValue interface{}) interface{} {
	if ptr == nil || reflect.ValueOf(ptr).IsNil() {
		return defaultValue
	}
	return reflect.ValueOf(ptr).Elem().Interface()
}
