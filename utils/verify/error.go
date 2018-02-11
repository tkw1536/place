package verify

import (
	"fmt"
	"strings"
)

// ValidationError is an error that failed to validate a given field
type ValidationError interface {
	Error() string
	Messages() []FieldError
}

type verr struct {
	fields []FieldError
}

func (ve verr) Error() string {
	var messages []string
	for _, fe := range ve.fields {
		messages = append(messages, fe.Error())
	}
	return strings.Join(messages, "\n")
}
func (ve verr) Messages() []FieldError { return ve.fields }

// NewValidationError creates a new validation error from a list of field errors
func NewValidationError(fields []FieldError) ValidationError {
	return verr{fields}
}

// FieldError represents a validation error of a single field
type FieldError interface {
	Field() string
	Error() string
}

type ferr struct {
	field string
	msg   string
}

func (fe ferr) Field() string { return fe.field }
func (fe ferr) Error() string { return fe.msg }

// FmtFieldError formates a new field error
func FmtFieldError(field string, format string, a ...interface{}) FieldError {
	return ferr{
		field,
		fmt.Sprintf(format, a...),
	}
}

// AsFieldError turns an error into a fieldError
func AsFieldError(field string, err error) FieldError {
	return ferr{
		field,
		err.Error(),
	}
}
