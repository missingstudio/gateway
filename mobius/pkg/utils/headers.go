package utils

import (
	"net/http"
	"reflect"
)

// UnmarshalHeader unmarshals an http.Header into a struct
func UnmarshalHeader(header http.Header, v interface{}) error {
	// Iterate over the fields in the struct
	for i := 0; i < reflect.TypeOf(v).Elem().NumField(); i++ {
		field := reflect.TypeOf(v).Elem().Field(i)
		tag := field.Tag.Get("json") // Get the tag value

		// If the tag is not empty, try to get the corresponding value from the header
		if tag != "" {
			value := header.Get(tag)
			// Set the value in the struct field
			reflect.ValueOf(v).Elem().Field(i).SetString(value)
		}
	}

	return nil
}
