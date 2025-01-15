package hash

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"

	"github.com/mansio-gmbh/goapiutils/must"
)

func SHA256[T any](t T) (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(data)
	return base64.StdEncoding.EncodeToString(hash[:]), nil
}

func MustSHA256[T any](t T) string {
	return must.Must(SHA256(t))
}

func SHA512[T any](t T) (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	hash := sha512.Sum512(data)
	return base64.StdEncoding.EncodeToString(hash[:]), nil
}

func MustSHA512[T any](t T) string {
	return must.Must(SHA512(t))
}
