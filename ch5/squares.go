// squares 返回一个函数，后者包含下一次要用到的平方根

package main
import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x*x
	}
}

func main() {
	f := squares()
	fmt.Println(f())  // "1"
	fmt.Println(f())  // "3"
	fmt.Println(f())  // "9"
}