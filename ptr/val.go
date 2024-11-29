package ptr

func Val[T any](t T) *T {
	return &t
}
