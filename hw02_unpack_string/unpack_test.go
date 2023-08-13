package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "п2д", expected: "ппд"},
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "llov", expected: "llov"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "aaab0", expected: "aaa"},
	}

	runTest(t, tests)
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackNonASCII(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "ллов", expected: "ллов"},

		{input: "ф2ф3ин3", expected: "фффффиннн"},
		{input: "ллов", expected: "ллов"},
		{input: "абв0г", expected: "абг"},
		{input: "п2", expected: "пп"},
	}

	runTest(t, tests)
}

func runTest(t *testing.T, tests []struct {
	input    string
	expected string
}) {
	t.Helper()
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}
