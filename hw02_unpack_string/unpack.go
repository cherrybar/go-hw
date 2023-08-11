package main

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func isDigit(char string) (int, bool) {
	val, err := strconv.Atoi(char)
	return val, err != nil
}

func removeLastLetterFromBuffer(buffer *bytes.Buffer) {
	strFromBuffer := buffer.String()
	currStr := strFromBuffer[0 : len(strFromBuffer)-1]
	var bufferTemp bytes.Buffer
	bufferTemp.WriteString(currStr)
	*buffer = bufferTemp
}

func Unpack(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	if _, isNan := isDigit(string(str[0])); !isNan {
		return "", ErrInvalidString
	}

	var buffer bytes.Buffer
	for index, char := range str {
		letter := string(char)

		if counter, isNan := isDigit(letter); isNan {
			buffer.WriteString(letter)
		} else {
			lastLetter := string(str[index-1])

			if _, isNan := isDigit(lastLetter); !isNan {
				return "", ErrInvalidString
			}

			if counter == 0 {
				removeLastLetterFromBuffer(&buffer)
			} else {
				buffer.WriteString(strings.Repeat(lastLetter, counter-1))
			}
		}
	}

	return buffer.String(), nil
}
