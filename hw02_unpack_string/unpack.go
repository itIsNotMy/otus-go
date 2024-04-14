package hw02unpackstring

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if valid, err := validator(s); !valid {
		return "", err
	}

	return string(unpacking(s)), nil
}

func unpacking(s string) []rune {
	result := make([]rune, 0, len(s))

	if len(s) <= 1 {
		return []rune(s)
	}

	r := []rune(s)

	for i := len(r) - 1; i >= 0; i-- {
		if isNumeric(r[i]) {
			if int(r[i]) == 48 {
				i--
				continue
			}

			for j := 48; j < int(r[i])-1; j++ {
				if res, val := runeKeyExist(i-2, r); res && isBackslash(val) {
					result = append(result, r[i-1], r[i-2])
				} else {
					result = append(result, r[i-1])
				}
			}
		} else {
			result = append(result, r[i])
		}
	}

	for i, j := 0, len(result)-1; i < len(result)/2; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func runeKeyExist(k int, r []rune) (bool, rune) {
	var val rune

	for key, val := range r {
		if k == key {
			return true, val
		}
	}

	return false, val
}

func isBackslash(s int32) bool {
	return s == 92
}

func isNumeric(s int32) bool {
	if 48 <= s && s <= 57 {
		return true
	}

	return false
}

func numbersExist(s string) bool {
	r := []rune(s)

	for i := 0; i < len(s)-1; i++ {
		if isNumeric(r[i]) && isNumeric(r[i+1]) {
			return true
		}
	}

	return false
}

func firstIsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	r := []rune(s)

	return isNumeric(r[0])
}

func validator(s string) (bool, error) {
	if firstIsNumeric(s) || numbersExist(s) {
		return false, ErrInvalidString
	}

	return true, nil
}
