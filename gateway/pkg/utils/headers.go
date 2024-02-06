package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/missingstudio/studio/backend/config"
	"github.com/missingstudio/studio/common/errors"
)

var ErrGatewayConfigHeaderNotValid = errors.New(fmt.Errorf("x-ms-config header is not valid"))

func isJSON(s string, v interface{}) bool {
	return json.Unmarshal([]byte(s), v) == nil
}

func UnmarshalConfigHeaders(header http.Header, v interface{}) error {
	msconfig := header.Get(config.XMSConfig)
	if msconfig == "" && isJSON(msconfig, v) {
		return ErrGatewayConfigHeaderNotValid
	}
	return nil
}

// UnmarshalHeader unmarshals an http.Header into a struct
func UnmarshalHeader(header http.Header, v interface{}) error {
	fields := reflect.ValueOf(v).Elem()

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Type().Field(i)
		headerKey := field.Tag.Get("json")
		defaultValue := field.Tag.Get("default")

		if headerValue := header.Get(headerKey); headerValue != "" {
			setFieldValue(fields.Field(i), headerValue)
		} else if defaultValue != "" {
			setFieldValue(fields.Field(i), defaultValue)
		}
	}

	return nil
}

func setFieldValue(field reflect.Value, value string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			field.SetInt(intValue)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err == nil {
			field.SetUint(uintValue)
		}
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err == nil {
			field.SetFloat(floatValue)
		}
	}
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
