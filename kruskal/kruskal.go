package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	From   int
	To     int
	Weight int
}

func Find(v int, parent []int) int {
	if v == parent[v] {
		return v
	}

	parent[v] = Find(parent[v], parent)

	return parent[v]
}

func Union(x, y int, parent []int) {
	rootX := Find(x, parent)
	rootY := Find(y, parent)
	if rootX != rootY {
		parent[rootX] = rootY
	}
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

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	parent := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		parent[i] = i
	}

	allWeight := 0
	for _, e := range edges {
		rootFrom := Find(e.From, parent)
		rootTo := Find(e.To, parent)

		if rootFrom != rootTo {
			allWeight += e.Weight
			Union(rootFrom, rootTo, parent)
		}
	}
}
