package main

import (
	"fmt"
	"os"
	"sort"
)

// Определение структуры ребра
type Edge struct {
	start, end, weight int
}

// Структура для системы непересекающихся множеств (DSU)
type DisjointSet struct {
	parent, rank []int
}

// Создание нового DSU
func NewDisjointSet(vertexCount int) *DisjointSet {
	parent := make([]int, vertexCount)
	rank := make([]int, vertexCount)
	for i := range parent {
		parent[i] = i
	}
	return &DisjointSet{parent, rank}
}

// Поиск сжатия пути
func (set *DisjointSet) Find(v int) int {
	if set.parent[v] != v {
		set.parent[v] = set.Find(set.parent[v])
	}
	return set.parent[v]
}

// Объединение множеств
func (set *DisjointSet) Union(x, y int) {
	rootX := set.Find(x)
	rootY := set.Find(y)
	if rootX != rootY {
		if set.rank[rootX] > set.rank[rootY] {
			set.parent[rootY] = rootX
		} else if set.rank[rootX] < set.rank[rootY] {
			set.parent[rootX] = rootY
		} else {
			set.parent[rootY] = rootX
			set.rank[rootX]++
		}
	}
}

// Алгоритм Крускала для нахождения MST
func KruskalMST(vertices int, edges []Edge) []Edge {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	mst := make([]Edge, 0)
	set := NewDisjointSet(vertices)

	for _, e := range edges {
		if set.Find(e.start) != set.Find(e.end) {
			set.Union(e.start, e.end)
			mst = append(mst, e)
		}
	}

	return mst
}

// Функция для генерации файла в формате DOT
func generateDOT(vertices int, edges, mst []Edge, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	file.WriteString("graph MST {\n")
	file.WriteString("    node [shape=circle, style=filled, fillcolor=lightgray];\n")

	// Записываем все рёбра графа (обычные — серые, MST — зелёные)
	for _, e := range edges {
		color := "gray"
		for _, m := range mst {
			if (m.start == e.start && m.end == e.end) || (m.start == e.end && m.end == e.start) {
				color = "green"
				break
			}
		}
		file.WriteString(fmt.Sprintf("    %d -- %d [label=%d, color=%s];\n", e.start, e.end, e.weight, color))
	}

	file.WriteString("}\n")
	fmt.Println("Файл с графом сохранён:", filename)
}

func main() {
	edges := []Edge{
		{0, 1, 10}, {0, 2, 6}, {0, 3, 5},
		{1, 3, 15}, {2, 3, 4},
	}

	mst := KruskalMST(4, edges)

	fmt.Println("Edges in MST:")
	for _, e := range mst {
		fmt.Printf("%d - %d: %d\n", e.start, e.end, e.weight)
	}

	// Генерация файла для визуализации
	generateDOT(4, edges, mst, "graph.dot")
}
