package ptr

func Val[T any](t T) *T {
	return &t
}

func ValOrNil[T comparable](t T) *T {
	var defaultT T
	if t == defaultT {
		return nil
	}
	return &t
}

type OrDefaultOptFn[T any] func() T

func WithDefaultValue[T any](t T) func() T {
	return func() T {
		return t
	}
}

func OrDefault[T any](t *T, opts ...OrDefaultOptFn[T]) T {
	if t == nil {
		if len(opts) > 0 {
			return opts[0]()
		}

		var defaultT T
		return defaultT
	}
	return *t
}
