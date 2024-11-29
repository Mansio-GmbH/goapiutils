package equals

type Equaler[T any] interface {
	IsEqual(other *T) bool
}

func PtrEq[T Equaler[T]](a, b *T) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return false
	}
	return (*a).IsEqual(b)
}
func ArrEq[T Equaler[T]](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if !t.IsEqual(&b[i]) {
			return false
		}
	}
	return true
}

func Ptr[T comparable](a, b *T) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return false
	}
	return *a == *b
}

func Arr[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if t != b[i] {
			return false
		}
	}
	return true
}
