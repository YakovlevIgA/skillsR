package main

import (
	"container/heap"
	"fmt"
)

// Структура для представления ребра
type edge struct {
	start, end, cost int
}

// Структура для хранения графа (список смежности)
type Graph map[int][]edge

// Очередь с приоритетами (минимальная куча)
type PriorityQueue []edge

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(edge))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// Функция алгоритма Прима
func prim(graph Graph, start int) []edge {
	mst := []edge{}
	visited := make(map[int]bool)
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Добавляем стартовую вершину в очередь
	visited[start] = true
	for _, e := range graph[start] {
		heap.Push(pq, e)
	}

	// Пока есть рёбра в очереди
	for pq.Len() > 0 {
		minEdge := heap.Pop(pq).(edge)
		if visited[minEdge.end] {
			continue
		}

		// Добавляем ребро в MST
		visited[minEdge.end] = true
		mst = append(mst, minEdge)

		// Добавляем все рёбра новой вершины
		for _, e := range graph[minEdge.end] {
			if !visited[e.end] {
				heap.Push(pq, e)
			}
		}
	}

	return mst
}

func main() {
	graph := Graph{
		0: {{0, 1, 10}, {0, 2, 6}, {0, 3, 5}},
		1: {{1, 0, 10}, {1, 3, 15}},
		2: {{2, 0, 6}, {2, 3, 4}},
		3: {{3, 0, 5}, {3, 1, 15}, {3, 2, 4}},
	}

	startNode := 0
	mst := prim(graph, startNode)

	fmt.Println("Рёбра в MST:")
	for _, e := range mst {
		fmt.Printf("%d - %d (вес %d)\n", e.start, e.end, e.cost)
	}
}
