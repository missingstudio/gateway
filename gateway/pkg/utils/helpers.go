package utils

import (
	"crypto/rand"
	"encoding/hex"
	"reflect"
)

const API_KEY_LENGTH = 40

func GetDefaultValue(ptr interface{}, defaultValue interface{}) interface{} {
	if ptr == nil || reflect.ValueOf(ptr).IsNil() {
		return defaultValue
	}
	return reflect.ValueOf(ptr).Elem().Interface()
}

// GenerateSecureAPIKey should generate a secure string to be used as an API key.
func GenerateSecureAPIKey() (string, error) {
	// Each byte of the binary data as two hexadecimal characters
	bytes := make([]byte, API_KEY_LENGTH/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return "ais-" + hex.EncodeToString(bytes), nil
}

func MaskString(input string) string {
	// Define the masking character
	maskChar := '*'

	// Convert the string to a slice of runes for easy manipulation
	runes := []rune(input)

	// Determine the number of characters to mask
	maskStart := 6
	maskEnd := len(runes) - 6

	for i := maskStart; i < maskEnd; i++ {
		runes[i] = maskChar
	}

	return string(runes)
}
