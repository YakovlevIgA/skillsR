package main

import (
	"fmt"
	"math"
)

type Edge struct {
	to, cost int
}

type Graph struct {
	vertices int
	edges    [][]Edge
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make([][]Edge, vertices),
	}
}

func (g *Graph) AddEdge(from, to, cost int) {
	g.edges[from] = append(g.edges[from], Edge{to, cost})
}

func (g *Graph) Dijkstra(source int) ([]int, [][]int) {
	minDist := make([]int, g.vertices)
	prev := make([]int, g.vertices)
	visited := make([]bool, g.vertices)

	for i := range minDist {
		minDist[i] = math.MaxInt32
		prev[i] = -1
	}

	minDist[source] = 0

	for i := 0; i < g.vertices; i++ {
		u := -1
		for v := 0; v < g.vertices; v++ {
			if !visited[v] && (u == -1 || minDist[v] < minDist[u]) {
				u = v
			}
		}

		if minDist[u] == math.MaxInt32 {
			break
		}

		visited[u] = true
		for _, e := range g.edges[u] {
			if minDist[u]+e.cost < minDist[e.to] {
				minDist[e.to] = minDist[u] + e.cost
				prev[e.to] = u
			}
		}
	}

	// Восстанавливаем пути
	paths := make([][]int, g.vertices)
	for i := range paths {
		paths[i] = getPath(i, prev)
	}

	return minDist, paths
}

// Восстанавливает путь от source до node
func getPath(node int, prev []int) []int {
	path := []int{}
	for node != -1 {
		path = append([]int{node}, path...)
		node = prev[node]
	}
	return path
}

func main() {
	graph := NewGraph(5)
	graph.AddEdge(0, 1, 10)
	graph.AddEdge(0, 2, 3)
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 2)
	graph.AddEdge(2, 1, 4)
	graph.AddEdge(2, 3, 8)
	graph.AddEdge(2, 4, 2)
	graph.AddEdge(3, 4, 7)
	graph.AddEdge(4, 3, 9)

	source := 0
	distances, paths := graph.Dijkstra(source)

	fmt.Println("Кратчайшие расстояния от вершины", source)
	for i, d := range distances {
		fmt.Printf("До %d: %d, путь: %v\n", i, d, paths[i])
	}
}
