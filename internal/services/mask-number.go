package services

import (
	"strings"
	"unicode"
)

// MaskPhoneNumber masks all digits except the last 2 digits of the phone number.
func MaskPhoneNumber(phoneNumber string) string {
	var masked strings.Builder
	n := len(phoneNumber)

	// Iterate through the phone number
	for i, char := range phoneNumber {
		if unicode.IsDigit(char) {

			// Mask all but the last 2 digits
			if i < n-2 {
				masked.WriteRune('*')
			} else {
				masked.WriteRune(char)
			}
		} else {
			// Append non-digit characters as is
			masked.WriteRune(char)
		}
	}

	return masked.String()
}
