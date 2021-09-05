package main

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// slice 仍有增长空间，扩展 slice 内容
		z = x[:zlen]
	} else {
		// slice 已无空间，为它分配一个新的底层数组
		// 为了达到分摊线性复杂性，容量扩展一倍
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)  // 内置 copy 函数
	}
	copy(z[len(x):], y)
	return z
}