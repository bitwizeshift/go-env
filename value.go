package env

import (
	"reflect"
	"time"
)

// Value represents an environment variable value, and provides basic utilities
// for converting to commonly used types.
type Value string

// Unmarshal the value into the given type.
//
// See [Unmarshal] for more details on what can be returned from this function.
func (v Value) Unmarshal(value any, opts ...UnmarshalOption) error {
	if value == nil {
		return nil
	}

	const key = "Value"
	tag := &tagOptions{
		key:   key,
		value: string(v),
		set:   true,
		sep:   ",",
	}
	for _, opt := range opts {
		opt.apply(tag)
	}
	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		rv = rv.Elem()
	}

	return decodeValue(nil, tag, key, rv.Type(), rv, nil)
}

// String returns the value as a string.
func (v Value) String() string {
	return string(v)
}

// Bool returns the value as a bool and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Bool() (bool, error) {
	var result bool
	err := v.Unmarshal(&result)
	return result, err
}

// Int returns the value as an int and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Int() (int, error) {
	var result int
	err := v.Unmarshal(&result)
	return result, err
}

// Int8 returns the value as an int8 and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Int8() (int8, error) {
	var result int8
	err := v.Unmarshal(&result)
	return result, err
}

// Int16 returns the value as an int16 and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Int16() (int16, error) {
	var result int16
	err := v.Unmarshal(&result)
	return result, err
}

// Int32 returns the value as an int32 and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Int32() (int32, error) {
	var result int32
	err := v.Unmarshal(&result)
	return result, err
}

// Int64 returns the value as an int64 and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Int64() (int64, error) {
	var result int64
	err := v.Unmarshal(&result)
	return result, err
}

// Uint returns the value as an uint and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Uint() (uint, error) {
	var result uint
	err := v.Unmarshal(&result)
	return result, err
}

// Uint8 returns the value as an uint8 and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Uint8() (uint8, error) {
	var result uint8
	err := v.Unmarshal(&result)
	return result, err
}

// Uint16 returns the value as an uint16 and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Uint16() (uint16, error) {
	var result uint16
	err := v.Unmarshal(&result)
	return result, err
}

// Uint32 returns the value as an uint32 and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Uint32() (uint32, error) {
	var result uint32
	err := v.Unmarshal(&result)
	return result, err
}

// Uint64 returns the value as an uint64 and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Uint64() (uint64, error) {
	var result uint64
	err := v.Unmarshal(&result)
	return result, err
}

// Float32 returns the value as a float32 and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Float32() (float32, error) {
	var result float32
	err := v.Unmarshal(&result)
	return result, err
}

// Float64 returns the value as a float64 and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Float64() (float64, error) {
	var result float64
	err := v.Unmarshal(&result)
	return result, err
}

// Duration returns the value as a [time.Duration] and returns any errors that
// may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Duration() (time.Duration, error) {
	var result time.Duration
	err := v.Unmarshal(&result)
	return result, err
}

// Time returns the value as a [time.Time] and returns any errors that may occur.
// See [Unmarshal] for more details on the possible errors that may be returned.
func (v Value) Time() (time.Time, error) {
	var result time.Time
	err := v.Unmarshal(&result)
	return result, err
}
