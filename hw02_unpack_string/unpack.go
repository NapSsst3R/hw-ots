package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inString string) (string, error) {
	var stringBuilder strings.Builder
	var symbolRepeat string
	var symbolEscaped bool

	for _, symbolRune := range inString {
		currentSymbol := string(symbolRune)
		switch {
		case symbolEscaped:
			if !(unicode.IsDigit(symbolRune) || currentSymbol == `\`) {
				return "", ErrInvalidString
			}
			symbolRepeat = currentSymbol
			symbolEscaped = false

		case currentSymbol == `\`:
			stringBuilder.WriteString(symbolRepeat)
			symbolRepeat = ""
			symbolEscaped = true

		case unicode.IsDigit(symbolRune):
			if symbolRepeat == "" {
				return "", ErrInvalidString
			}
			repeatCount, err := strconv.Atoi(currentSymbol)
			if err != nil {
				return "", err
			}
			stringBuilder.WriteString(strings.Repeat(symbolRepeat, repeatCount))
			symbolRepeat = ""

		default:
			stringBuilder.WriteString(symbolRepeat)
			symbolRepeat = currentSymbol
		}
	}
	if symbolEscaped {
		return "", ErrInvalidString
	}
	stringBuilder.WriteString(symbolRepeat)
	return stringBuilder.String(), nil
}
