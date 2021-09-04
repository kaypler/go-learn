// basename 简化版
package main
func basename(s string) string {
	slash := strings.LastIndex(s, "/") // 如果没找到“/”，则 slash 取值 -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}