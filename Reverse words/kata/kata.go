package kata

import (
	"fmt"
	"strings"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	result += " "
	return
}

func ReverseWords(in string) (out string) {
	splitted := strings.Split(in, " ")
	for _, v := range splitted {
		out += Reverse(v)
	}
	return
}

// func main() {
// 	fmt.Println(ReverseWords("hello world nice day to die"))
// }
