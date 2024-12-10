/*
Package env is a small package for managing environment variables.
*/
package env

import "fmt"

// Environment is a map of environment variables.
type Environment map[string]string

// Marshal converts a value into an environment map.
func Marshal(any) (Environment, error) {
	return nil, fmt.Errorf("not implemented")
}
