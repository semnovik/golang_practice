package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var newStr string
	strLen := utf8.RuneCountInString(str)

	for i := 0; i <= strLen; i++ {
		fmt.Println(str[i])
	}

	return newStr, nil
}

func main() {
	var str = "a4bc2d5e"
	var newStr = ""
	var cache = ""
	strLen := utf8.RuneCountInString(str)

	for i := 0; i < strLen; i++ {

		if unicode.IsDigit(rune(str[i])) == true {
			num, _ := strconv.Atoi(string(str[i]))
			fmt.Println("Тут цифра", num)
			newStr += strings.Repeat(cache, num-1)
		} else {
			cache = string(str[i])
			newStr += cache
			fmt.Println(newStr)
		}
	}

}
