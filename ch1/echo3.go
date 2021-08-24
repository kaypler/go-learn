// 输出命令行参数
// 如果有大量的数据需要处理，前两种的代价会比较大
// 一个简单高效的方式是使用 strings 包中的 Join 函数
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
