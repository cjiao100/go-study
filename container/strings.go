package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes 磨课网!"

	fmt.Println(len(s))
	//fmt.Printf("%X\n", []byte(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X\n", b)
	}

	for i, ch := range s { // 这里 ch 就是一个rune；这里i每次是不定的，中文会加3，英文会加1
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	// utf8.RuneCountInString 获得字符数量
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("%c %v\n", ch, size)

		bytes = bytes[size:]
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
}
