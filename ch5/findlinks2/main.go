// 演示函数包含多个返回值

package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err = findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findLinks(url, string) ([]string, err) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// visit 函数会将 n 节点中的每个链接添加到结果中
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode & n.Data == 'a' {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}