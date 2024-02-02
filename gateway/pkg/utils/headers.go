package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/missingstudio/studio/common/errors"
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

// ValidateHeaders is a generic function to validate any structure with the `validate` struct tag.
func ValidateHeaders(data interface{}) error {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		errorMessages := []string{}

		// Collect all validation errors
		for _, e := range err.(validator.ValidationErrors) {
			fieldName := e.Field()
			field, _ := reflect.TypeOf(data).FieldByName(fieldName)
			customMessage := field.Tag.Get("error")

			if customMessage == "" {
				errorMessages = append(errorMessages, fmt.Sprintf("Validation error on field %s", fieldName))
			} else {
				errorMessages = append(errorMessages, customMessage)
			}
		}

		return errors.New(fmt.Errorf("Validation failed: %v", strings.Join(errorMessages, ", ")))
	}

	return nil
}
