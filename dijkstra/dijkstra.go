package main

import (
	"algorithms/structures"
	"container/heap"
	"fmt"
)

type Edge struct {
	To     int
	Weight int
}

type Item struct {
	node int
	dist int
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	graph := make(map[int][]Edge, n+1)

	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)

		e := Edge{To: v, Weight: w}

		graph[u] = append(graph[u], e)
	}

	dist := make([]int, n+1)
	dist[1] = 0
	for i := 2; i < n+1; i++ {
		dist[i] = 1e9
	}

	visited := make([]bool, n+1)

	pq := make(structures.IntHeap, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: 1, dist: 0})

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Item).node

		if visited[u] {
			continue
		}

		visited[u] = true

		for _, edge := range graph[u] {
			v := edge.To
			weight := edge.Weight

			if dist[v] > dist[u]+weight {
				dist[v] = dist[u] + weight
				heap.Push(&pq, &Item{node: v, dist: dist[v]})
			}
		}
	}
}
