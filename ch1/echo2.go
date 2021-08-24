// 输出命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// 每一次迭代，range 产生一对值：索引和这个索引处的值
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
