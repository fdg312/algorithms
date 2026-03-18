package main

import (
	"fmt"
	"sort"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	graph := make(map[int][]int, n+1)
	indegree := make([]int, n+1)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
		indegree[u] = indegree[u] + 1
		indegree[v] = indegree[v] + 1
	}

	nodes := make([]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = i + 1
	}

	sort.Slice(nodes, func(i, j int) bool {
		return indegree[nodes[i]] > indegree[nodes[j]]
	})

	colors := make([]int, n+1)
	for i := range colors {
		colors[i] = -1
	}

	notAvaliable := make([]bool, n+1)

	colors[1] = 1

	for _, i := range nodes {
		for _, u := range graph[i] {
			if colors[u] != -1 {
				notAvaliable[colors[u]] = true
			}
		}
		for j := 1; j < n+1; j++ {
			if !notAvaliable[j] {
				colors[i] = j
				break
			}
		}
		for _, u := range graph[i] {
			if colors[u] != -1 {
				notAvaliable[colors[u]] = false
			}
		}
	}
}
