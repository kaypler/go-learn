// str 逐个处理 Unicode 字符
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))                     // "13"
	fmt.Println(utf8.RuneCountInString(s))  // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// []rune 可以返回字符串的 Unicode 码点序列
	r := []rune(s)
	fmt.Printf("% x\n", s)
	fmt.Printf("%x\n", r)
}
