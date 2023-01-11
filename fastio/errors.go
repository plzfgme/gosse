package fastio

import "errors"

var (
	// ErrKeySize is returned when key size isn't equal to ClientKeySize
	ErrKeySize = errors.New("Key size error")
	// ErrIDTooLong is returned when id is too long
	ErrIDTooLong = errors.New("ID too long")
)
