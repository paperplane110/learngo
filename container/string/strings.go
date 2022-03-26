package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Yes哈哈哈哈哈！"
	fmt.Println(len(s))
	fmt.Printf("%s\n", s)

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}

	fmt.Println()
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("%c %d\n", ch, size)
		bytes = bytes[size:]
	}

	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}

	fmt.Println()
	fmt.Println(strings.Fields("hhhh kdjfslkdfj   jsdkfjd      1"))
	fmt.Println(strings.Split("1321,2323,4343", ","))
	fmt.Println(strings.Join([]string{"12312", "23232", "3123"}, ","))
	fmt.Println(strings.Contains("Hello world", "world"))
	fmt.Println(strings.Index("Hello world", "world"))
}
