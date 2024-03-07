package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/missingstudio/ai/common/errors"
)

var ErrRouterConfigHeaderNotValid = errors.New(fmt.Errorf("x-ms-config header is not valid"))

// ValidateHeaders is a generic function to validate any structure with the `validate` struct tag.
func ValidateHeaders(data any) error {
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
