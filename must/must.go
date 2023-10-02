package must

import (
	"log"
)

func Must[T any](t T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func Must2[T any, U any](t T, u U, err error) (T, U) {
	if err != nil {
		log.Fatal(err)
	}
	return t, u
}

func Must3[T any, U any, V any](t T, u U, v V, err error) (T, U, V) {
	if err != nil {
		log.Fatal(err)
	}
	return t, u, v
}

func WithoutError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
