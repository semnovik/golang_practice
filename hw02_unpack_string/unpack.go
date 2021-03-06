package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	got := []rune(str)
	newStr := ""
	cachedSymbol := ""

	for i := 0; i < len(got); i++ {
		symbol := got[i]

		switch {
		// Проверяем, что первый символ не цифра
		case i == 0 && unicode.IsDigit(symbol):
			return "", ErrInvalidString
		// Проверяем, что нет двухзначных цифр
		case unicode.IsDigit(symbol) && unicode.IsDigit(got[i-1]):
			return "", ErrInvalidString
		// Если в строке содержится 0
		case symbol == 48:
			newStr = strings.Replace(newStr, cachedSymbol, "", 1)
		// Если цифра после буквы
		case unicode.IsDigit(symbol) && symbol != 48:
			num, err := strconv.Atoi(string(symbol))
			if err != nil {
				return "", err
			}
			newStr += strings.Repeat(cachedSymbol, num-1)
		// Если буква
		case !unicode.IsDigit(symbol):
			cachedSymbol = string(symbol)
			newStr += cachedSymbol
		}
	}
	return newStr, nil
}
