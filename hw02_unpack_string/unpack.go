package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var builder strings.Builder

	var previous rune

	for key, val := range s {
		if isNumeric(val) {
			if key == 0 {
				return "", ErrInvalidString
			}

			if isNumeric(previous) {
				return "", ErrInvalidString
			}

			if isNull(val) {
				str := builder.String()
				updatedStr := str[:len(str)-1]
				builder.Reset()
				builder.WriteString(updatedStr)
			}

			for i := '0'; i < val-1; i++ {
				builder.WriteRune(previous)
			}
		} else {
			builder.WriteRune(val)
		}

		previous = val
	}

	return builder.String(), nil
}

func isNull(s int32) bool {
	return s == '0'
}

func isNumeric(s int32) bool {
	return '0' <= s && s <= '9'
}
