package shared

import (
	"fmt"
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
