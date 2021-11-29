package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	const emptyRune = rune(0)
	unpackedString := strings.Builder{}
	lastRune := emptyRune
	isEscaped := false
	for _, currentRune := range packedString {

		if isEscaped {
			if unicode.IsLetter(currentRune) {
				// впрочем, что плохого в экранировании буквы? ;)
				return "", ErrInvalidString
			}
			if lastRune != emptyRune {
				unpackedString.WriteRune(lastRune)
			}
			lastRune = currentRune
			isEscaped = false
			continue
		}

		if currentRune == '\\' {
			isEscaped = true
			continue
		}

		if unicode.IsLetter(currentRune) && lastRune == emptyRune {
			lastRune = currentRune
			continue
		}

		if unicode.IsLetter(currentRune) {
			unpackedString.WriteRune(lastRune)
			lastRune = currentRune
			continue
		}

		if unicode.IsDigit(currentRune) && lastRune == emptyRune {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(currentRune) {
			count, _ := strconv.Atoi(string(currentRune))
			for i := 0; i < count; i++ {
				unpackedString.WriteRune(lastRune)
			}
			lastRune = emptyRune
			continue
		}

		return "", ErrInvalidString
	}

	if isEscaped {
		return "", ErrInvalidString
	}

	if lastRune == emptyRune {
		return unpackedString.String(), nil
	}

	unpackedString.WriteRune(lastRune)

	return unpackedString.String(), nil
}
