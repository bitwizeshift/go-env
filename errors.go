package env

import (
	"fmt"
	"reflect"
)

// InvalidTagOptionError is an error that occurs when an invalid tag option is
// used in a struct field tag.
type InvalidTagOptionError struct {
	// Key is the environment variable key that caused the error.
	Key string

	// Option is the tag option that caused the error.
	Option string

	// Type is the type of the field the error occurred on.
	Type reflect.Type

	// Field is the struct field that caused the error. This is nil if the type
	// is not a struct field.
	Field *reflect.StructField
}

func (e *InvalidTagOptionError) Error() string {
	if e.Field == nil {
		return fmt.Sprintf("env: invalid tag option '%s' for env variable '%s'", e.Option, e.Key)
	}
	return fmt.Sprintf("env: invalid tag option '%s' on field '%s'", e.Option, e.Field.Name)
}

var _ error = (*InvalidTagOptionError)(nil)

// InvalidTypeError is an error that occurs when an unsupported type is used.
// This can occur when a struct field is not a pointer or when a struct field
// is not a supported type.
type InvalidTypeError struct {
	// Key is the environment variable key assigned to the type.
	Key string

	// Type is the type that caused the error.
	Type reflect.Type

	// Field is the struct field that caused the error. This is nil if the type
	// is not a struct field.
	Field *reflect.StructField
}

func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf("env: invalid type '%s' for env variable '%s'", e.Type, e.Key)
}

var _ error = (*InvalidTypeError)(nil)

// ParseError is an error that occurs when a value cannot be parsed from an
// environment variable.
type ParseError struct {
	// Key is the environment variable key that caused the error.
	Key string

	// Value is the value that caused the error.
	Value string

	// Type is the type that caused the error.
	Type reflect.Type

	// Err is the underlying error that was triggered during parsing.
	Err error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("env: unable to parse %s from env variable %s: %v", e.Key, e.Type, e.Err)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

var _ error = (*ParseError)(nil)

// RequirementError is an error that occurs when a required environment variable
// is missing.
type RequirementError struct {
	Key  string
	Type reflect.Type
}

func (e *RequirementError) Error() string {
	return fmt.Sprintf("env: missing required env value '%s'", e.Key)
}

var _ error = (*RequirementError)(nil)
