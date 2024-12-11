package env_test

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"rodusek.dev/pkg/env"
)

type Custom int

func (c *Custom) UnmarshalEnv(b []byte) error {
	v, err := strconv.ParseInt(string(b), 10, 0)
	*c = Custom(v)
	return err
}

type CustomText int

func (c *CustomText) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 0)
	*c = CustomText(v)
	return err
}

type OptionalEnv struct {
	PtrString       *string         `env:"PTR_STRING"`
	String          string          `env:"STRING"`
	Bool            bool            `env:"BOOL"`
	Int8            int8            `env:"INT8"`
	Int16           int16           `env:"INT16"`
	Int32           int32           `env:"INT32"`
	Int64           int64           `env:"INT64"`
	Int             int             `env:"INT"`
	Uint8           uint8           `env:"UINT8"`
	Uint16          uint16          `env:"UINT16"`
	Uint32          uint32          `env:"UINT32"`
	Uint64          uint64          `env:"UINT64"`
	Uint            uint            `env:"UINT"`
	Float32         float32         `env:"FLOAT32"`
	Float64         float64         `env:"FLOAT64"`
	Duration        time.Duration   `env:"DURATION"`
	Time            time.Time       `env:"TIME"`
	StringSlice     []string        `env:"STRING_SLICE,sep=;"`
	DurationSlice   []time.Duration `env:"DURATION_SLICE"`
	Unmarshaler     Custom          `env:"UNMARSHALER"`
	PtrUnmarshaler  *Custom         `env:"PTR_UNMARSHALER"`
	TextUnmarshaler CustomText      `env:"TEXT_UNMARSHALER"`
	Pointers        ***int          `env:"POINTERS"`
	AnonymousInt    int
}

func TestUnmarshal_OptionalKeys_ParsesValues(t *testing.T) {
	testCases := []struct {
		name        string
		want        any
		environment string
	}{
		// Optional Pointer String
		{
			name: "Optional Pointer String",
			want: &OptionalEnv{
				PtrString: func() *string { s := "Hello World"; return &s }(),
			},
			environment: "PTR_STRING=Hello World",
		},
		// Optional String
		{
			name: "Optional String",
			want: &OptionalEnv{
				String: "Hello World",
			},
			environment: "STRING=Hello World",
		},
		// Optional Int8
		{
			name: "Optional decimal Int8",
			want: &OptionalEnv{
				Int8: 42,
			},
			environment: "INT8=42",
		}, {
			name: "Optional hex Int8",
			want: &OptionalEnv{
				Int8: 0x0F,
			},
			environment: "INT8=0x0f",
		}, {
			name: "Optional octal Int8",
			want: &OptionalEnv{
				Int8: 0o12,
			},
			environment: "INT8=0o12",
		}, {
			name: "Optional binary Int8",
			want: &OptionalEnv{
				Int8: 0b1101101,
			},
			environment: "INT8=0b1101101",
		},
		// Optional Int16
		{
			name: "Optional decimal Int16",
			want: &OptionalEnv{
				Int16: 256,
			},
			environment: "INT16=256",
		}, {
			name: "Optional hex Int16",
			want: &OptionalEnv{
				Int16: 0x0FFF,
			},
			environment: "INT16=0x0fff",
		}, {
			name: "Optional octal Int16",
			want: &OptionalEnv{
				Int16: 0o1234,
			},
			environment: "INT16=0o1234",
		}, {
			name: "Optional binary Int16",
			want: &OptionalEnv{
				Int16: 0b1101101_00000000,
			},
			environment: "INT16=0b110110100000000",
		},

		// Optional Int32
		{
			name: "Optional decimal Int32",
			want: &OptionalEnv{
				Int32: 75535,
			},
			environment: "INT32=75535",
		}, {
			name: "Optional hex Int32",
			want: &OptionalEnv{
				Int32: 0x0FFFFFFF,
			},
			environment: "INT32=0x0fffffff",
		}, {
			name: "Optional octal Int32",
			want: &OptionalEnv{
				Int32: 0o1234567,
			},
			environment: "INT32=0o1234567",
		}, {
			name: "Optional binary Int32",
			want: &OptionalEnv{
				Int32: 0b1101101_00000000_11111111_00000000,
			},
			environment: "INT32=0b1101101000000001111111100000000",
		},

		// Optional Int64
		{
			name: "Optional decimal Int64",
			want: &OptionalEnv{
				Int64: 75535,
			},
			environment: "INT64=75535",
		}, {
			name: "Optional hex Int64",
			want: &OptionalEnv{
				Int64: 0x0FFFFFFF,
			},
			environment: "INT64=0x0fffffff",
		}, {
			name: "Optional octal Int64",
			want: &OptionalEnv{
				Int64: 0o1234567,
			},
			environment: "INT64=0o1234567",
		}, {
			name: "Optional binary Int64",
			want: &OptionalEnv{
				Int64: 0b1101101_00000000_11111111_00000000,
			},
			environment: "INT64=0b1101101000000001111111100000000",
		},

		// Optional Int
		{
			name: "Optional decimal Int",
			want: &OptionalEnv{
				Int: 75535,
			},
			environment: "INT=75535",
		}, {
			name: "Optional hex Int",
			want: &OptionalEnv{
				Int: 0x0FFFFFFF,
			},
			environment: "INT=0x0fffffff",
		}, {
			name: "Optional octal Int",
			want: &OptionalEnv{
				Int: 0o1234567,
			},
			environment: "INT=0o1234567",
		}, {
			name: "Optional binary Int",
			want: &OptionalEnv{
				Int: 0b1101101_00000000_11111111_00000000,
			},
			environment: "INT=0b1101101000000001111111100000000",
		},

		// Optional Uint8
		{
			name: "Optional Uint8",
			want: &OptionalEnv{
				Uint8: 42,
			},
			environment: "UINT8=42",
		},

		// Optional Uint16
		{
			name: "Optional Uint16",
			want: &OptionalEnv{
				Uint16: 0xffff,
			},
			environment: "UINT16=0xffff",
		},

		// Optional Uint32
		{
			name: "Optional Uint32",
			want: &OptionalEnv{
				Uint32: 0xffff_ffff,
			},
			environment: "UINT32=0xffffffff",
		},

		// Optional Uint64
		{
			name: "Optional Uint64",
			want: &OptionalEnv{
				Uint64: 0xffffffff_ffffffff,
			},
			environment: "UINT64=0xffffffffffffffff",
		},

		// Optional Uint
		{
			name: "Optional Uint",
			want: &OptionalEnv{
				Uint: 0xffff_ffff,
			},
			environment: "UINT=0xffffffff",
		},

		// Optional Float32
		{
			name: "Optional Float32",
			want: &OptionalEnv{
				Float32: 3.14,
			},
			environment: "FLOAT32=3.14",
		},

		// Optional Float64
		{
			name: "Optional Float64",
			want: &OptionalEnv{
				Float64: 3.14,
			},
			environment: "FLOAT64=3.14",
		},

		// Optional Duration
		{
			name: "Optional seconds Duration",
			want: &OptionalEnv{
				Duration: 5 * time.Second,
			},
			environment: "DURATION=5s",
		}, {
			name: "Optional minutes Duration",
			want: &OptionalEnv{
				Duration: 5 * time.Minute,
			},
			environment: "DURATION=5m",
		}, {
			name: "Optional minutes Duration",
			want: &OptionalEnv{
				Duration: 5 * time.Hour,
			},
			environment: "DURATION=5h",
		},

		// Optional Bool
		{
			name: "Optional bool numeric",
			want: &OptionalEnv{
				Bool: true,
			},
			environment: "BOOL=1",
		}, {
			name: "Optional bool word",
			want: &OptionalEnv{
				Bool: true,
			},
			environment: "BOOL=true",
		},

		// Optional time
		{
			name: "Optional time",
			want: &OptionalEnv{
				Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			environment: "TIME=2021-01-01T00:00:00Z",
		},

		// Unnamed Int Field
		{
			name: "Unnamed Int Field",
			want: &OptionalEnv{
				AnonymousInt: 42,
			},
			environment: "ANONYMOUS_INT=42",
		},

		// Slice
		{
			name: "String Slice",
			want: &OptionalEnv{
				StringSlice: []string{"Hello", "World"},
			},
			environment: "STRING_SLICE=Hello;World",
		},
		{
			name: "Duration Slice",
			want: &OptionalEnv{
				DurationSlice: []time.Duration{5 * time.Second, 5 * time.Minute, 5 * time.Hour},
			},
			environment: "DURATION_SLICE=5s,5m,5h",
		},

		// Unmarshaler
		{
			name: "Unmarshaler",
			want: &OptionalEnv{
				Unmarshaler: Custom(42),
			},
			environment: "UNMARSHALER=42",
		}, {
			name: "Pointer Unmarshaler",
			want: &OptionalEnv{
				PtrUnmarshaler: func() *Custom { c := Custom(42); return &c }(),
			},
			environment: "PTR_UNMARSHALER=42",
		}, {
			name: "TextUnmarshaler",
			want: &OptionalEnv{
				TextUnmarshaler: CustomText(42),
			},
			environment: "TEXT_UNMARSHALER=42",
		},
		// Pointers
		{
			name: "Pointers",
			want: &OptionalEnv{
				Pointers: func() ***int {
					i := 42
					p := &i
					pp := &p
					return &pp
				}(),
			},
			environment: "POINTERS=42",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setenv(t, tc.environment)

			var out OptionalEnv
			err := env.Unmarshal(&out)
			if err != nil {
				t.Fatalf("Unmarshal(%s): unexpected error: %v", tc.name, err)
			}

			if got, want := &out, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Unmarshal(%s): got '%v', want '%v'", tc.name, got, want)
			}
		})
	}
}

func setenv(t *testing.T, str string, args ...any) {
	lines := strings.Split(fmt.Sprintf(str, args...), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		key := parts[0]
		value := ""
		if len(parts) == 2 {
			value = parts[1]
		}
		t.Setenv(key, value)
	}
}

func TestUnmarshal_RequiredKeyNotSet_ReturnsError(t *testing.T) {
	type RequiredEnv struct {
		Required string `env:"REQUIRED,required"`
	}

	var out RequiredEnv
	err := env.Unmarshal(&out)

	var requiredErr *env.RequirementError
	if !errors.As(err, &requiredErr) {
		t.Fatalf("Unmarshal(): expected RequirementError, got %T", err)
	}

	if requiredErr.Key != "REQUIRED" {
		t.Errorf("Unmarshal(): expected RequirementError, got %T", err)
	}
}

func TestUnmarshal_RequiredKeySet_ParsesValues(t *testing.T) {
	type RequiredEnv struct {
		Required string `env:"REQUIRED,required"`
	}

	want := "Hello World"
	setenv(t, "REQUIRED=%v", want)

	var out RequiredEnv
	err := env.Unmarshal(&out)
	if err != nil {
		t.Fatalf("Unmarshal(): unexpected error: %v", err)
	}

	if got := out.Required; got != want {
		t.Errorf("Unmarshal(): got '%s', want '%s'", got, want)
	}
}
