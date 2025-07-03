package job

import (
	"log/slog"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {

	return emailRegex.MatchString(email)

}

func IsStrongPassword(password string) bool {

	var hasUpper, hasNum, hasChar bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasNum = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasChar = true
		}
	}

	return hasChar && hasNum && hasUpper
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("error hashing the password")
		return "", err
	}
	return string(hashedBytes), nil
}