package validations

import (
	"fmt"
	"regexp"

	"expenze-io.com/internal/body"
	"github.com/nyaruka/phonenumbers"
)

type ValidateRegister struct {
	body *body.RegistrationBody
}

// constructor
func New(registrationBody *body.RegistrationBody) *ValidateRegister {
	return &ValidateRegister{
		body: registrationBody,
	}
}

// validate phonenumber
func validatePhonenumber(mobilenumber, phonecode string) *ValidateError {
	phonenumber := fmt.Sprintf("+%s%s", phonecode, mobilenumber)

	// Parse the phone number with the country code
	parsedNumber, err := phonenumbers.Parse(phonenumber, "")
	if err != nil {
		error := fmt.Sprintf("Error parsing phone number: %v", err)

		return &ValidateError{
			Key:   "mobilenumber",
			Error: error,
			Value: mobilenumber,
		}
	}

	isValid := phonenumbers.IsValidNumber(parsedNumber)

	if !isValid {
		return &ValidateError{
			Key:   "mobilenumber",
			Error: "Mobile number is invalid",
			Value: mobilenumber,
		}
	}

	return nil
}

func MinMaxValidation(field MinMaxValidationFields) *ValidateError {
	if field.Min == nil {
		defaultMin := 0
		field.Min = &defaultMin
	}

	length := len(field.FieldValue)

	if length < *field.Min {
		return &ValidateError{
			Key:   field.FieldName,
			Error: fmt.Sprintf("%s must be at least %d character long", field.FieldName, *field.Min),
			Value: field.FieldValue,
		}
	}

	if field.Max == nil {
		defaultMax := 0
		field.Max = &defaultMax
	}

	if *field.Max > 0 {
		if length > *field.Max {
			return &ValidateError{
				Key:   field.FieldName,
				Error: fmt.Sprintf("%s much be at most %d character long", field.FieldName, *field.Max),
				Value: field.FieldValue,
			}
		}

	}

	return nil
}

// custom validation function for password strength
func ValidatePassword(password string) *ValidateError {
	if len(password) < 8 {
		return &ValidateError{
			Key:   "password",
			Error: "Password should be more then 8 character",
			Value: password,
		}
	}

	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	if !hasLowercase {
		return &ValidateError{
			Key:   "password",
			Error: "Password should have at least 1 lower character",
			Value: password,
		}
	}

	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)

	if !hasUppercase {
		return &ValidateError{
			Key:   "password",
			Error: "Password should have at least 1 upper character",
			Value: password,
		}
	}

	hasSpecialChar := regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString(password)

	if !hasSpecialChar {
		return &ValidateError{
			Key:   "password",
			Error: "Password should be at least 1 special character",
			Value: password,
		}
	}

	hasSpace := regexp.MustCompile(`\s`).MatchString(password)

	if hasSpace {
		return &ValidateError{
			Key:   "password",
			Error: "Password should not have space",
			Value: password,
		}
	}

	return nil
}

// custom validation function for email
func ValidateEmail(email string) *ValidateError {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	validEmail := regexp.MustCompile(emailRegex).MatchString(email)

	if !validEmail {
		return &ValidateError{
			Key: "email",

			Error: "Email should be valid",
			Value: email,
		}
	}

	return nil
}

// validate registration body
func (v *ValidateRegister) ValidateRegistration() *ValidateError {
	// Validate password
	if err := ValidatePassword(v.body.Password); err != nil {
		return err
	}

	// Validate email
	if err := ValidateEmail(v.body.EmailID); err != nil {
		return err
	}

	// Validate firstname
	if err := MinMaxValidation(MinMaxValidationFields{
		Min:        IntPtr(4),
		FieldName:  "firstname",
		FieldValue: v.body.Firstname,
	}); err != nil {
		return err
	}

	// Validate lastname
	if err := MinMaxValidation(MinMaxValidationFields{
		Min:        IntPtr(3),
		FieldName:  "lastname",
		FieldValue: v.body.Lastname,
	}); err != nil {
		return err
	}

	// validate mobile number
	if err := validatePhonenumber(v.body.MobilieNumber, v.body.PhoneCode); err != nil {
		return err
	}

	return nil
}
