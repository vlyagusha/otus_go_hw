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
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "ггг0д", expected: "ггд"},
		{input: "жжж0з0", expected: "жж"},
		{input: "й0", expected: ""},
		{input: "л2", expected: "лл"},
		{input: "ф1ы1ц1", expected: "фыц"},
		{input: "\n", expected: "\n"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "\n\t", expected: "\n\t"},
		{input: "й\t5уфы", expected: "й\t\t\t\t\tуфы"},
		// uncomment if task with asterisk completed
		{input: `\\`, expected: `\`},
		{input: `\\\\`, expected: `\\`},
		{input: `\4`, expected: `4`},
		{input: `\4\5`, expected: `45`},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{
		"3abc",
		"45",
		"aaa10b",
		"3абв",
		"ййй10ц",
		"ф56",
		"ф56ы",
		"ф5ы67",
		"0",
		".",
		"&ы4",
		"ы4?",
		`\\\`,
		`\\\\\`,
		`\ф`,
		`ф\ф`,
		`\5\ы`,
		`\n`,
		`\t`,
	}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
