package getutflength

import (
	"errors"
	"unicode/utf8"
)

var ErrInvalidUTF8 error = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}
	return utf8.RuneCount(input), nil
}
