package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	reversedString := "Hello, OTUS!"

	fmt.Println(stringutil.Reverse(reversedString))
}
