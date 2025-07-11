package dto

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	uppercaseRegex   = regexp.MustCompile(`[A-Z]`)
	lowercaseRegex   = regexp.MustCompile(`[a-z]`)
	digitRegex       = regexp.MustCompile(`[0-9]`)
	specialCharRegex = regexp.MustCompile(`[\W_]`)
)

func StrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	length := len(password)

	return length >= 8 && length <= 12 &&
		uppercaseRegex.MatchString(password) &&
		lowercaseRegex.MatchString(password) &&
		digitRegex.MatchString(password) &&
		specialCharRegex.MatchString(password)
}

type LoginDto struct {
	Name     string `json:"name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	PASSWORD string `json:"password" binding:"required,strongpassword"`
}
