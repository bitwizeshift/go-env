package env

// UnmarshalOption is an option that can be passed to the [Unmarshal] or
// [Environment.Unmarshal] functions.
type UnmarshalOption interface {
	apply(*tagOptions)
}

type apply func(*tagOptions)

func (a apply) apply(tag *tagOptions) {
	a(tag)
}

// Separator returns an [UnmarshalOption] that sets the default separator for
// splitting values for slice values.
//
// This is the _only_ way to set a custom separator when using [Value]'s
// unmarshal functionality, since values are not part of a struct and therefore
// cannot provide the `env` sep tag.
func Separator(sep string) UnmarshalOption {
	return apply(func(tag *tagOptions) {
		tag.sep = sep
	})
}
