// rev 就地反转一个整型 slice 中的元素
package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// 反转整个数组 a
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"

	// 向左移动两个元素
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
}

