package main

import "fmt"

var a = b + c // 最后把 a 初始化为 3
var b = f()   // 通过调用 f 接着把 b 初始化为 2
var c = 1     // 首先初始化为 1

func f() int { return c + 1 }

func main() {
	fmt.Printf("a=%v, b=%v, c=%v", a, b, c)
}
