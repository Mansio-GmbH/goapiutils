package ptr

func MapNilIfEmpty[K comparable, T any](m map[K]T) map[K]T {
	if len(m) == 0 {
		return nil
	}
	return m
}

func NotNil[T any](v *T) bool {
	return v != nil
}

func IsNil[T any](v *T) bool {
	return v == nil
}
