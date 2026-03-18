package main

import (
	"algorithms/structures"
	"fmt"
	"slices"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	graph := make(map[int][]int, n+1)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n+1)
	visited[1] = true

	dist := make([]int, n+1)
	dist[1] = 0

	q := structures.NewQueue[int]()
	q.Enqueue(1)

	maxID := 1

	for q.Size() > 0 {
		s, _ := q.Peek()
		q.Dequeue()

		for _, v := range graph[s] {
			if visited[v] {
				continue
			}
			dist[v] = dist[s] + 1
			if dist[v] > dist[maxID] {
				maxID = v
			}
			visited[v] = true
			q.Enqueue(v)
		}
	}

	visited = make([]bool, n+1)
	visited[maxID] = true

	dist = make([]int, n+1)
	dist[maxID] = 0

	q = structures.NewQueue[int]()
	q.Enqueue(maxID)

	for q.Size() > 0 {
		s, _ := q.Peek()
		q.Dequeue()

		for _, v := range graph[s] {
			if visited[v] {
				continue
			}
			dist[v] = dist[s] + 1
			visited[v] = true
			q.Enqueue(v)
		}
	}

	fmt.Println(slices.Max(dist))
}
