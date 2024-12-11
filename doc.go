/*
Package env is a small package for managing environment variables.

Unlike the standard [os.Getenv] approach, this package models itself after
the [encoding/json] package, providing a way to marshal and unmarshal
environment variables into a structured format using env tags.
*/
package env
