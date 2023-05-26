package main
import (
	"fmt"
	"strings"
	"unicode"
)

func StringChallenge(str string) string {
	length := len(str)
	if length <= 7 || length >= 31 {
		return "false"
	}

	if strings.Contains(str, "password") {
		return "false"
	}

	var capital, number, symbol bool

	for _, c := range str {
		if unicode.IsUpper(c) {
			capital = true
		}
		if unicode.IsNumber(c) {
			number = true
		}
		if unicode.IsMark(c) {
			symbol = true
		}
	}

	if capital && number && symbol {
		return "true"
	}

	return "false"
}

func main() {

	fmt.Println(StringChallenge("passWord123!!!!"))
}