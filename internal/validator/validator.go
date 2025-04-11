package validator

import "regexp"

// EmailRX will parse the regualr expression for an email address
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Validator defines a map with fields that will contain errors
type Validator struct {
	FieldErrors map[string]string
}

// New instantiates a new Validator struct
func New() *Validator {
	return &Validator{
		FieldErrors: make(map[string]string),
	}
}

// IsValid returns wether the validator contains any errors
func (v *Validator) IsValid() bool {
	return len(v.FieldErrors) == 0
}

// AddError will append an new error entry to the Validator map
func (v *Validator) AddError(key string, value string) {
	_, exists := v.FieldErrors[key]
	if !exists {
		v.FieldErrors[key] = value
	}
}

// Check verifies if a condition is true and, if so, adds a new Error entry to the Validator by calling AddError
func (v *Validator) Check(condition bool, key, value string) {
	if condition {
		v.AddError(key, value)
	}
}

// IsEmptyString returns if as string is empty
func IsEmptyString(str string) bool {
	return str == ""
}

// WithinMaxLimit returns if a string has less characters than a specified maxLimit
func WithinMaxLimit(str string, maxLimit int) bool {
	return len(str) <= maxLimit
}

// WithMinLimit returns if a string has more characters than a specified minLimite
func WithinMinLimit(str string, minLimit int) bool {
	return len(str) >= minLimit
}

// IsValidEmail returns whether a email addres is valid
func IsValidEmail(str string) bool {
	return EmailRX.MatchString(str)
}
