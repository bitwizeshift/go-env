package env

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

// Environment is a map of environment variables.
//
// This map is a simple key-value store where the key is the environment
// variable name and the value is the environment variable value. It may be
// used to unmarshal values without requiring the real environment to be
// modified, such as through the dotenv sub-package.
//
// This type is not thread-safe. If you need to write to the access while
// concurrently reading it, you should use a mutex to protect it.
type Environment map[string]Value

// Load the current environment variables into a new [Environment] instance.
//
// The returned map will contain all the elements returned from [os.Environ].
func Load() Environment {
	keys := os.Environ()
	env := make(Environment, len(keys))
	for _, key := range keys {
		parts := strings.SplitN(key, "=", 2)
		env[parts[0]] = Value(parts[1])
	}
	return env
}

// New creates a new empty environment.
func New() Environment {
	return make(Environment)
}

// Get the value of the environment variable with the given key, falling back
// to the real environment as if by using [os.Getenv].
//
// Note: To search this without falling back to [os.Getenv], use the map
// access notation instead:
//
//	value := env[key]
func (e Environment) Get(key string) Value {
	val, _ := e.Lookup(key)
	return val
}

// Lookup the value of the environment variable with the given key.
//
// If the environment variable does not exist, it will be looked up in the
// real environment using [os.LookupEnv]. If it still does not exist, the
// second return value will be false.
//
// Note: To search this without falling back to [os.LookupEnv], use the map
// access notation instead:
//
//	value, ok := env[key]
func (e Environment) Lookup(key string) (value Value, ok bool) {
	if e != nil {
		value, ok = e[key]
	}
	if !ok {
		valueStr, ok := os.LookupEnv(key)
		if !ok {
			return "", false
		}
		value = Value(valueStr)
	}
	return
}

// Set the value of the environment variable with the given key.
func (e *Environment) Set(key string, value Value) {
	if *e == nil {
		*e = make(Environment)
	}
	(*e)[key] = value
}

// Unset the environment variable with the given key.
func (e Environment) Unset(key string) {
	delete(e, key)
}

// Contains returns true if the environment variable with the given key exists.
func (e Environment) Contains(key string) bool {
	if e == nil {
		return false
	}
	_, ok := e[key]
	if !ok {
		_, ok = os.LookupEnv(key)
	}
	return ok
}

// Export sets the environment variables in the current process.
func (e Environment) Export() {
	for key, value := range e {
		os.Setenv(key, string(value))
	}
}

// ExportCmd sets the environment variables into the specified subprocess
// command object.
func (e Environment) ExportCmd(cmd *exec.Cmd) {
	for key, value := range e {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%v", key, value))
	}
}

// Unmarshal the environment variables into the given struct.
// See the documentation for [Unmarshal] for more details on what can be
// returned from this function.
func (e Environment) Unmarshal(out any, opts ...UnmarshalOption) error {
	rv := reflect.ValueOf(out)
	lookup := func(key string) (string, bool) {
		value, ok := e.Lookup(key)
		return string(value), ok
	}
	return decode(lookup, rv, opts...)
}
