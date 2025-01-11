package ptr

func ArrNilIfEmpty[T any](a []T) []T {
	if len(a) == 0 {
		return nil
	}
	return a
}
