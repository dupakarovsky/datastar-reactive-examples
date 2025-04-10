package validator

import "regexp"

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	FieldErrors map[string]string
}

func New() *Validator {
	return &Validator{
		FieldErrors: make(map[string]string),
	}
}

func (v *Validator) IsValid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) AddError(key string, value string) {
	_, exists := v.FieldErrors[key]
	if !exists {
		v.FieldErrors[key] = value
	}
}

func (v *Validator) Check(condition bool, key, value string) {
	if condition {
		v.AddError(key, value)
	}
}

func IsEmptyString(str string) bool {
	return str == ""
}

func WithinMaxLimit(str string, maxLimit int) bool {
	return len(str) <= maxLimit
}

func WithinMinLimit(str string, minLimit int) bool {
	return len(str) >= minLimit
}

func IsValidEmail(str string) bool {
	return EmailRX.MatchString(str)
}
