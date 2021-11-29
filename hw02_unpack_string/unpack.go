package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const emptyRune = rune(0)

func Unpack(packedString string) (string, error) {
	unpackedString := strings.Builder{}
	lastRune := emptyRune

	for _, currentRune := range packedString {
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

	if unicode.IsLetter(lastRune) {
		unpackedString.WriteRune(lastRune)
		return unpackedString.String(), nil
	}

	if lastRune == emptyRune {
		return unpackedString.String(), nil
	}

	return "", ErrInvalidString
}
