// outline 递归遍历所有 HTML 文本中的节点树，并输出树的结构
// ./fetch https://golang.org | .outline
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main()  {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // 把标签压入栈
		fmt.Println(stack)
	}
	for c := c.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}