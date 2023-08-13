package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	runes := []rune(str)

	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	builder := strings.Builder{}
	var prevLetter rune

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			if unicode.IsDigit(runes[i-1]) {
				return "", ErrInvalidString
			}

			if counter, _ := strconv.Atoi(string(runes[i])); counter != 0 {
				runeArray := []rune(strings.Repeat(string(prevLetter), counter))

				for _, run := range runeArray {
					builder.WriteRune(run)
				}

			}
			prevLetter = 0

			continue
		}

		if prevLetter != 0 {
			builder.WriteRune(prevLetter)
		}

		if i == len(runes)-1 {
			builder.WriteRune(runes[i])
		}

		prevLetter = runes[i]
	}

	return builder.String(), nil
}
