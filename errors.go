package env

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	// ErrEnv is a generic error that occurs when an error is related to the
	// environment package.
	errEnv = errors.New("env")

	// ErrRequirement is an error that occurs when a required environment variable
	// is missing. When an error is determined to be this type, it can be converted
	// into a [RequirementError].
	ErrRequirement = fmt.Errorf("%w: requirement error", errEnv)

	// ErrInvalidTagOption is an error that occurs when an invalid tag option is
	// used in a struct field tag. When an error is determined to be this type, it
	// can be converted into an [InvalidTagOptionError].
	ErrInvalidTagOption = fmt.Errorf("%w: invalid tag option", errEnv)

	// ErrInvalidType is an error that occurs when an unsupported type is used.
	// This can occur when a struct field is not a pointer or when a struct field
	// is not a supported type. When an error is determined to be this type, it can
	// be converted into an [InvalidTypeError].
	ErrInvalidType = fmt.Errorf("%w: invalid type", errEnv)

	// ErrParse is an error that occurs when a value cannot be parsed from an
	// environment variable. When an error is determined to be this type, it can
	// be converted into a [ParseError].
	ErrParse = fmt.Errorf("%w: parse error", errEnv)
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

func (e *InvalidTagOptionError) Unwrap() error {
	return ErrInvalidTagOption
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
	return fmt.Sprintf("%v '%s' for env variable '%s'", ErrInvalidType, e.Type, e.Key)
}

func (e *InvalidTypeError) Unwrap() error {
	return ErrInvalidType
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

func (e *ParseError) Unwrap() []error {
	return []error{e.Err, ErrParse}
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

func (e *RequirementError) Unwrap() error {
	return ErrRequirement
}

var _ error = (*RequirementError)(nil)
