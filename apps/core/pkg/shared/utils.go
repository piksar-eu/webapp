package shared

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func SanitizeEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		return "", fmt.Errorf("incorrect email")
	}

	return email, nil
}

func MapToStruct(data interface{}, target interface{}) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonBytes, target)
}
func StrictUnmarshal[T any](data []byte) *T {
	var target T

	if err := json.Unmarshal(data, &target); err != nil {
		return nil
	}

	marshaled, err := json.Marshal(target)
	if err != nil {
		return nil
	}

	if JsonEqual(data, marshaled) {
		return &target
	}

	return nil
}

func JsonEqual(a, b []byte) bool {
	var j1, j2 interface{}
	if err := json.Unmarshal(a, &j1); err != nil {
		return false
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false
	}
	return reflect.DeepEqual(j1, j2)
}
