package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	graph := make(map[int][]int, n)

	for i := 0; i < n; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
}
