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
	var isPrevDigit bool

	for _, char := range str {
		if unicode.IsDigit(char) {
			if isPrevDigit {
				return "", ErrInvalidString
			}

			if counter, _ := strconv.Atoi(string(char)); counter != 0 {
				runeArray := []rune(strings.Repeat(string(prevLetter), counter))

				for _, run := range runeArray {
					builder.WriteRune(run)
				}
			}
			prevLetter = 0
			isPrevDigit = true
			continue
		}
		isPrevDigit = false

		if prevLetter != 0 {
			builder.WriteRune(prevLetter)
		}

		prevLetter = char
	}

	if !unicode.IsDigit(prevLetter) && prevLetter != 0 {
		builder.WriteRune(prevLetter)
	}

	return builder.String(), nil
}
