package env

import (
	"encoding"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Unmarshaler is an interface that allows for custom unmarshaling of
// environment variables.
type Unmarshaler interface {
	// UnmarshalEnv unmarshals the provided value from an environment string into
	// the implementing type.
	UnmarshalEnv(value []byte) error
}

// Unmarshal reads values from the current environment and parses values into
// the provided output struct.
//
// Fields on the output struct are interpreted based on `env` tag options which
// defines the variable key to read from, and any additional options.
// If this tag is not set, the field name is converted to screaming
// snake case and used instead (e.g. the field `ProjectName` would use the
// environment variable `PROJECT_NAME`). Unexported fields are ignored.
//
// A nil `out` parameter is valid and will return nil without error.
//
// This function supports parsing values from the environment for the following
// types:
//
//   - string types
//   - integral types (byte, int, int8, int16, int32, int64, uint, uint8,
//     uint16, uint32, uint64)
//   - floating point types (float32, float64)
//   - boolean types
//   - [time.Duration] (using [time.ParseDuration] format)
//   - [time.Time] (using [time.Parse], using all common time format layouts)
//   - [Unmarshaler]
//   - [encoding.TextUnmarshaler]
//   - slices of any of the above supported types
//
// This makes use of the `env` tag to specify the environment variable key to
// read from.
//
// Fields may be marked as required by adding the `required` option to the tag.
// Slices may have custom separators (default is ',') that may be specified with
// the `sep` option. For example:
//
//	type Environment struct {
//		ProjectName string        `env:"PROJECT_NAME,required"`
//		Timeout     time.Duration `env:"TIMEOUT"`
//		Path        []string      `env:"PATH,required,sep=;"`
//	}
//
// On error, this function may return one of the following error types:
//
//   - [RequirementError] when a required environment variable was not defined.
//   - [ParseError] when a value cannot be parsed from an environment variable.
//   - [InvalidTypeError] when an unsupported type is used without defining it
//     as a [Marshaler] or [encoding.TextUnmarshaler].
//   - [InvalidTagOptionError] when an invalid/unsupported tag option is used.
func Unmarshal(out any, opts ...UnmarshalOption) error {
	// Nothing in, no error taking it out. Seems reasonable?
	if out == nil {
		return nil
	}

	rv := reflect.ValueOf(out)
	return decode(os.LookupEnv, rv, opts...)
}

// lookup is a function that performs a string lookup on the environment.
// This is used internally to allow Unmarshal to be used with a custom env.
type lookup func(key string) (string, bool)

type tagOptions struct {
	key      string
	value    string
	set      bool
	required bool
	sep      string
}

func toScreamingSnake(s string) string {
	var builder strings.Builder
	prevLower := false
	for _, r := range s {
		if prevLower && unicode.IsUpper(r) {
			builder.WriteByte('_')
		}
		prevLower = unicode.IsLower(r)
		builder.WriteRune(r)
	}
	return strings.ToUpper(builder.String())
}

func readTag(lookup lookup, field *reflect.StructField, opts ...UnmarshalOption) (*tagOptions, error) {
	tag, ok := field.Tag.Lookup("env")
	if !ok {
		tag = toScreamingSnake(field.Name)
	}

	parts := strings.Split(tag, ",")
	key := parts[0]

	value, ok := lookup(key)
	tagOptions := &tagOptions{
		key:      key,
		value:    value,
		set:      ok,
		required: false,
		sep:      ",",
	}
	for _, opt := range opts {
		opt.apply(tagOptions)
	}
	for _, part := range parts[1:] {
		switch part {
		case "required":
			tagOptions.required = true
		default:
			if rest, ok := strings.CutPrefix(part, "sep="); ok {
				tagOptions.sep = rest
				continue
			}
			return nil, &InvalidTagOptionError{
				Key:    key,
				Option: part,
				Type:   field.Type,
				Field:  field,
			}
		}
	}
	return tagOptions, nil
}

func bitness(rt reflect.Type) int {
	switch rt.Kind() {
	case reflect.Int8, reflect.Uint8:
		return 8
	case reflect.Int16, reflect.Uint16:
		return 16
	case reflect.Int32, reflect.Uint32, reflect.Float32:
		return 32
	case reflect.Int64, reflect.Uint64, reflect.Float64:
		return 64
	case reflect.Int, reflect.Uint:
		return 0
	default:
		return rt.Bits()
	}
}

func decode(lookup lookup, rv reflect.Value, opts ...UnmarshalOption) error {
	rt := rv.Type()
	if rt.Kind() != reflect.Ptr {
		return fmt.Errorf("env: expected pointer, got '%s'", rt.String())
	}

	for rt.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return fmt.Errorf("env: cannot unmarshal into nil pointer")
		}
		rv = rv.Elem()
		rt = rt.Elem()
	}
	return decodeStruct(lookup, rv, rt, opts...)
}

func decodeStruct(lookup lookup, rv reflect.Value, rt reflect.Type, opts ...UnmarshalOption) error {
	if rt.Kind() != reflect.Struct {
		return &InvalidTypeError{
			Type: rt,
		}
	}

	length := rt.NumField()
	for i := 0; i < length; i++ {
		field := rt.Field(i)
		tag, err := readTag(lookup, &field, opts...)
		if err != nil {
			return err
		}

		if err := decodeValue(lookup, tag, field.Name, field.Type, rv.Field(i), &field); err != nil {
			return err
		}
	}
	return nil
}

var timeLayouts = []string{
	time.Layout,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
	time.DateTime,
	time.DateOnly,
	time.TimeOnly,
	time.Kitchen,
}

func pointsToStruct(rt reflect.Type) bool {
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	return rt.Kind() == reflect.Struct
}

func deref(rv reflect.Value, rt reflect.Type) (reflect.Value, reflect.Type) {
	for rt.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rt.Elem()))
		}
		rv = rv.Elem()
		rt = rt.Elem()
	}
	return rv, rt
}

func decodeValue(lookup lookup, tag *tagOptions, name string, rt reflect.Type, rv reflect.Value, field *reflect.StructField) error {
	if !rv.CanSet() {
		return fmt.Errorf("env: cannot set field '%s'", name)
	}

	if !tag.set {
		if tag.required {
			return &RequirementError{
				Key:  tag.key,
				Type: rt,
			}
		}
		return nil
	}

	rv, rt = deref(rv, rt)

	makeParseError := func(err error) error {
		errParse := ParseError{
			Key:   tag.key,
			Value: tag.value,
			Type:  rt,
			Err:   err,
		}
		return &errParse
	}

	// Try converting to Unmarshaler first
	if marshaler, ok := rv.Addr().Interface().(Unmarshaler); ok {
		if err := marshaler.UnmarshalEnv([]byte(tag.value)); err != nil {
			return makeParseError(err)
		}
	}

	// Fallback to TextUnmarshaler if it's available
	if marshaler, ok := rv.Addr().Interface().(encoding.TextUnmarshaler); ok {
		if err := marshaler.UnmarshalText([]byte(tag.value)); err != nil {
			return makeParseError(err)
		}
	}

	// Handle specific cases
	switch rt {
	case durationType:
		duration, err := time.ParseDuration(tag.value)
		if err != nil {
			return makeParseError(err)
		}
		rv.Set(reflect.ValueOf(duration))
		return nil
	case timeType:
		var err error
		for _, layout := range timeLayouts {
			var timeValue time.Time
			timeValue, err = time.Parse(layout, tag.value)
			if err != nil {
				continue
			}
			rv.Set(reflect.ValueOf(timeValue))
			err = nil
			break
		}
		if err != nil {
			return makeParseError(err)
		}
		return nil
	}

	// Handle decoding primitive types
	switch rt.Kind() {
	case reflect.String:
		rv.SetString(tag.value)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		integer, err := strconv.ParseInt(tag.value, 0, bitness(rt))
		if err != nil {
			return makeParseError(err)
		}
		rv.SetInt(integer)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		integer, err := strconv.ParseUint(tag.value, 0, bitness(rt))
		if err != nil {
			return makeParseError(err)
		}
		rv.SetUint(integer)
		return nil
	case reflect.Float32, reflect.Float64:
		value, err := strconv.ParseFloat(tag.value, bitness(rt))
		if err != nil {
			return makeParseError(err)
		}
		rv.SetFloat(value)
		return nil
	case reflect.Bool:
		value, err := strconv.ParseBool(tag.value)
		if err != nil {
			return makeParseError(err)
		}
		rv.SetBool(value)
		return nil
	case reflect.Slice:
		entries := strings.Split(tag.value, tag.sep)
		slice := reflect.MakeSlice(rt, 0, len(entries))
		for _, entry := range entries {
			elem := reflect.New(rt.Elem()).Elem()
			newTag := *tag
			newTag.value = entry
			if err := decodeValue(lookup, &newTag, name, rt.Elem(), elem, field); err != nil {
				return makeParseError(err)
			}
			slice = reflect.Append(slice, elem)
		}
		rv.Set(slice)
		return nil
	default:
		return &InvalidTypeError{
			Key:   tag.key,
			Type:  rt,
			Field: field,
		}
	}
}

var (
	durationType = reflect.TypeFor[time.Duration]()
	timeType     = reflect.TypeFor[time.Time]()
)

// Get retrieves the value of the environment variable with the given key and
// unmarshals it into the provided type. This is a strongly-typed equivalent
// of [os.Getenv].
//
// This function will only return errors if the environment variable is not set
// or if the value cannot be unmarshaled into the provided type correctly.
func Get[T any](name string) (got T, err error) {
	value, ok := os.LookupEnv(name)
	if !ok {
		err = &RequirementError{
			Key:  name,
			Type: reflect.TypeFor[T](),
		}
		return
	}
	err = Value(value).Decode(&got)
	return
}

// GetOr retrieves the value of the environment variable with the given key and
// unmarshals it into the provided type. If the environment variable is not set,
// the fallback value is returned instead. This is a strongly-typed equivalent
// of [os.Getenv] with a fallback.
//
// This function will only return errors if the value cannot be unmarshaled into
// the provided type correctly.
func GetOr[T any](name string, fallback T) (got T, err error) {
	value, ok := os.LookupEnv(name)
	if !ok {
		return fallback, nil
	}
	err = Value(value).Decode(&got)
	return
}
