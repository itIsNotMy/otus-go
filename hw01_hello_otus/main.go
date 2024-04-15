package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "a a a e e e s s s w w q k k k k l l l l l l x x u h h h h f "
	//
	//m := wordCount(s)
	//
	//fmt.Println(m)
	//fmt.Println(sortDesc(m))

	s := "d0w0r2l3kkk\nhgf\nngg"

	fmt.Println(Unpack(s))

}

func Unpack(s string) string {
	var builder strings.Builder

	//var previous rune

	r := []rune(s)

	for i := 0; i < len(r)+1; i++ {
		//if isNumeric(r[i]) {
		//	if i == 0 {
		//		return ""
		//	}
		//
		//	if isNumeric(previous) {
		//		return ""
		//	}
		//
		//	if isNull(r[i]) {
		//		continue
		//	}
		//
		//	for i := '0'; i < r[i]-1; i++ {
		//		builder.WriteRune(previous)
		//	}
		//	continue
		//}

		builder.WriteRune(r[i])

		//previous = r[i]
	}

	return builder.String()
}

func isNull(s int32) bool {
	return s == '0'
}

func isNumeric(s int32) bool {
	return '0' <= s && s <= '9'
}

//func wordCount(s string) map[string]int {
//	m := make(map[string]int)
//
//	var builder strings.Builder
//
//	for _, val := range s {
//
//		if isSpace(val) {
//			if _, ok := m[builder.String()]; ok {
//				m[builder.String()]++
//			} else {
//				m[builder.String()] = 1
//			}
//
//			builder.Reset()
//		} else {
//			builder.WriteRune(val)
//		}
//	}
//
//	return m
//}
//
//func sortDesc(m map[string]int) []string {
//	s := make([]string, len(m)+1, len(m)+1)
//
//	for k, v := range m {
//		for {
//			if s[v] != "" {
//				v = v + 1
//			} else {
//				s[v] = k
//				break
//			}
//		}
//	}
//
//	return s
//}
//
//func isSpace(i int32) bool {
//	return i == ' '
//}
