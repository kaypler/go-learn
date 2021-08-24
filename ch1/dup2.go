// 打印输入中多次出现的行的个数和文本
// 它从 stdin 或指定的文件列表读取
// 这个版本的 dup 使用“流式”模式读取输入，然后按需拆分为行，这样可以处理海量的输入
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// map 是一个使用 make 创建的数据结构的引用
// 当一个 map 被传递给一个函数时，函数接收到这个引用的副本
// 所以被调用函数中对于 map 数据结构中的改变对函数调用者使用的 map 引用也是可见的
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
