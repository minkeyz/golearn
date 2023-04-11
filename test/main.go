package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isDigit("0x12487115155"))
}

func isDigit(str string) bool {
	ss := str[2:]
	for _, x := range []rune(ss) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}
