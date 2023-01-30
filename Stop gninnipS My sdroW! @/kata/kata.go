package kata

import "strings"

func Spin(str []byte) string {
	for i, j2 := 0, len(str)-1; i < j2; i, j2 = i+1, j2-1 {
		str[i], str[j2] = str[j2], str[i]
	}
	return string(str)
	// var based []byte
	// j := len(str) - 1
	// for i := 0; i < len(str); i++ {
	// 	based[i] = str[j]
	// 	j--
	// }
	// return string(based)
}

func SpinWords(str string) string {
	splitted := strings.Split(str, " ")
	for _, v := range splitted {
		if len(v) >= 5 {
			v = Spin([]byte(v))
		}
	}
	return strings.Join(splitted, " ")
} // SpinWords
