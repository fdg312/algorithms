package main

import (
	"fmt"
)

type Edge struct {
	From   int
	To     int
	Weight int
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	edges := make([]Edge, 0, m)

	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)

		e := Edge{From: u, To: v, Weight: w}

		edges = append(edges, e)
	}

	dist := make([]int, n+1)
	dist[1] = 0
	for i := 2; i < n+1; i++ {
		dist[i] = 1e9
	}

	for i := 1; i < n; i++ {
		for _, e := range edges {
			if dist[e.From] < 1e9 && dist[e.To] > dist[e.From]+e.Weight {
				dist[e.To] = dist[e.From] + e.Weight
			}
		}
	}

	for _, e := range edges {
		if dist[e.From] < 1e9 && dist[e.To] > dist[e.From]+e.Weight {
			fmt.Println("Отрицательный цикл")
		}
	}
}
