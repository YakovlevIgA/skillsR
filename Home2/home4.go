package main

import "fmt"

func main() {
	m := make(map[string]int, 100) // в исходном случае мапа nil, здесь
	var a map[string]int
	fmt.Println(a)
	for _, word := range []string{"hello", "world", "from", "the", "best", "language",
		"in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
