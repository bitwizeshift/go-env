package env_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/env"
)

func TestValueBool(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    bool
		wantErr error
	}{
		{
			name:    "Valid boolean value",
			value:   env.Value("true"),
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Invalid boolean value",
			value:   env.Value("not_a_boolean"),
			want:    false,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Bool()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Bool(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Bool(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueInt(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    int
		wantErr error
	}{
		{
			name:    "Valid integer value",
			value:   env.Value("42"),
			want:    42,
			wantErr: nil,
		},
		{
			name:    "Invalid integer value",
			value:   env.Value("not_an_integer"),
			want:    0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Int()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Int(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Int(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueInt8(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    int8
		wantErr error
	}{
		{
			name:    "Valid int8 value",
			value:   env.Value("42"),
			want:    42,
			wantErr: nil,
		},
		{
			name:    "Invalid int8 value",
			value:   env.Value("not_an_int8"),
			want:    0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Int8()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Int8(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Int8(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueInt16(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    int16
		wantErr error
	}{
		{
			name:    "Valid int16 value",
			value:   env.Value("42"),
			want:    42,
			wantErr: nil,
		},
		{
			name:    "Invalid int16 value",
			value:   env.Value("not_an_int16"),
			want:    0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Int16()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Int16(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Int16(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueInt32(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    int32
		wantErr error
	}{
		{
			name:    "Valid int32 value",
			value:   env.Value("42"),
			want:    42,
			wantErr: nil,
		},
		{
			name:    "Invalid int32 value",
			value:   env.Value("not_an_int16"),
			want:    0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Int32()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Int32(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Int32(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueInt64(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    int64
		wantErr error
	}{
		{
			name:    "Valid int64 value",
			value:   env.Value("42"),
			want:    42,
			wantErr: nil,
		},
		{
			name:    "Invalid int64 value",
			value:   env.Value("not_an_int64"),
			want:    0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Int64()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Int64(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Int64(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueUint(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    uint
		wantErr error
	}{
		{
			name:    "Valid integer value",
			value:   env.Value("42"),
			want:    42,
			wantErr: nil,
		},
		{
			name:    "Invalid integer value",
			value:   env.Value("not_an_integer"),
			want:    0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Uint()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Uint(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Uint(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueUint8(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    uint8
		wantErr error
	}{
		{
			name:    "Valid uint8 value",
			value:   env.Value("42"),
			want:    42,
			wantErr: nil,
		},
		{
			name:    "Invalid uint8 value",
			value:   env.Value("not_an_uint8"),
			want:    0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Uint8()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Uint8(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Uint8(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueUint16(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    uint16
		wantErr error
	}{
		{
			name:    "Valid uint16 value",
			value:   env.Value("42"),
			want:    42,
			wantErr: nil,
		},
		{
			name:    "Invalid uint16 value",
			value:   env.Value("not_a_uint16"),
			want:    0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Uint16()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Uint16(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Uint16(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueUint32(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    uint32
		wantErr error
	}{
		{
			name:    "Valid uint32 value",
			value:   env.Value("42"),
			want:    42,
			wantErr: nil,
		},
		{
			name:    "Invalid uint32 value",
			value:   env.Value("not_a_uint32"),
			want:    0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Uint32()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Uint32(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Uint32(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueUint64(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    uint64
		wantErr error
	}{
		{
			name:    "Valid uint64 value",
			value:   env.Value("42"),
			want:    42,
			wantErr: nil,
		},
		{
			name:    "Invalid uint64 value",
			value:   env.Value("not_a_uint64"),
			want:    0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Uint64()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Uint64(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Uint64(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueFloat32(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    float32
		wantErr error
	}{
		{
			name:    "Valid float32 value",
			value:   env.Value("42.0"),
			want:    42.0,
			wantErr: nil,
		},
		{
			name:    "Invalid float32 value",
			value:   env.Value("not_a_float32"),
			want:    0.0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Float32()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Float32(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Float32(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueFloat64(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    float64
		wantErr error
	}{
		{
			name:    "Valid float64 value",
			value:   env.Value("42.0"),
			want:    42.0,
			wantErr: nil,
		},
		{
			name:    "Invalid float64 value",
			value:   env.Value("not_a_float64"),
			want:    0.0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Float64()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Float64(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Float64(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueDuration(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    time.Duration
		wantErr error
	}{
		{
			name:    "Valid duration value",
			value:   env.Value("5s"),
			want:    5 * time.Second,
			wantErr: nil,
		},
		{
			name:    "Invalid duration value",
			value:   env.Value("not_a_duration"),
			want:    0,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Duration()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Duration(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Duration(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueTime(t *testing.T) {
	testCases := []struct {
		name    string
		value   env.Value
		want    time.Time
		wantErr error
	}{
		{
			name:    "Valid time value",
			value:   env.Value("2021-01-01T00:00:00Z"),
			want:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			wantErr: nil,
		},
		{
			name:    "Invalid time value",
			value:   env.Value("not_a_time"),
			want:    time.Time{},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.value.Time()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Value.Time(%s): got error '%v', want error '%v'", tc.name, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Time(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func TestValueString(t *testing.T) {
	testCases := []struct {
		name  string
		value env.Value
		want  string
	}{
		{
			name:  "Valid string value",
			value: env.Value("hello"),
			want:  "hello",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.value.String()

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.String(%s): got '%v', want '%v'", tc.name, got, tc.want)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}

func TestValuePointer(t *testing.T) {
	testCases := []struct {
		name  string
		value env.Value
		want  *string
	}{
		{
			name:  "Valid string value",
			value: env.Value("hello"),
			want:  ptr("hello"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var got *string
			err := tc.value.Decode(&got)

			if err != nil {
				t.Fatalf("Value.Decode(%s): got error '%v', want error nil", tc.name, err)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Value.Decode(%s): got '%v', want '%v'", tc.name, *got, *tc.want)
			}
		})
	}
}
