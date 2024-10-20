package services

import (
	"regexp"
	"strings"
	"unicode"
)

// creating a struct
type MaskNumber struct {
	number *string
}

type MaskText struct {
	text *string
}

// creating a constructors
func NewMaskNumber(number *string) *MaskNumber {
	return &MaskNumber{
		number: number,
	}
}

func NewMaskText(text *string) *MaskText {
	return &MaskText{
		text: text,
	}
}

func (mn *MaskNumber) MaskNumber(lastVisibleDigit int) string {
	var masked strings.Builder
	n := len(*mn.number)

	// Iterate through the phone number
	for i, char := range *mn.number {
		if unicode.IsDigit(char) {
			// Mask all but the last digits based on the last visible digit
			if i < n-lastVisibleDigit {
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

func (mt *MaskText) MaskEmail() string {
	pattern := `(^[\w._%+-]+)(@[\w.-]+\.[a-zA-Z]{2,})`
	re := regexp.MustCompile(pattern)

	// masking email except domain
	maskedEmail := re.ReplaceAllString(*mt.text, "****$2")
	return maskedEmail
}
