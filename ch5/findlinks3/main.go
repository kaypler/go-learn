package main

func breadthFirst(f, func(item string) []string, worklist []string) {
	sean := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !sean[item] {
				sean[item] = true
				worklist = append(worklist, f(item)...) // "f(item)..."会把f返回的列表中的所有项添加到worklist中
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}