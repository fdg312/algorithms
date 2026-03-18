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

	dist := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = 1e9
			}
		}
	}

	for _, e := range edges {
		dist[e.From][e.To] = e.Weight
		// dist[e.To][e.From] = e.Weight // если граф неориентированный
	}

	for k := 1; k < n+1; k++ {
		for i := 1; i < n+1; i++ {
			for j := 1; j < n+1; j++ {
				if dist[i][k] != 1e9 && dist[k][j] != 1e9 && dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
}
