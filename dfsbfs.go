package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Graph структура для представления графа
type Graph map[int][]int

// addEdge добавляет рёбра в ненаправленный граф
func addEdge(graph Graph, from, to int) {
	graph[from] = append(graph[from], to)
	graph[to] = append(graph[to], from) // Для ненаправленного графа добавляем ребро в обе стороны
}

// dfs реализует алгоритм поиска в глубину
func dfs(graph Graph, start int) {
	visited := make(map[int]bool)
	var recur func(v int)
	recur = func(v int) {
		if visited[v] {
			return
		}
		visited[v] = true
		fmt.Println(v)
		for _, n := range graph[v] {
			recur(n)
		}
	}
	recur(start)
}

// bfs реализует алгоритм поиска в ширину
func bfs(graph Graph, start int) {
	queue := []int{start}
	visited := make(map[int]bool)
	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		fmt.Println(vertex)
		for _, n := range graph[vertex] {
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
			}
		}
	}
}

// readGraph читает граф из файла
func readGraph(filename string) (Graph, error) {
	graph := make(Graph)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			// Преобразуем строки в числа и добавляем рёбра
			from := stringToInt(parts[0])
			to := stringToInt(parts[1])
			addEdge(graph, from, to)
		}
	}
	return graph, scanner.Err()
}

// stringToInt преобразует строку в целое число
func stringToInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

func main() {

	filename := "graph.txt"
	graph, err := readGraph(filename)
	if err != nil {
		fmt.Println("Error reading graph:", err)
		return
	}

	fmt.Println("Graph structure:")
	fmt.Println(graph)

	fmt.Println("\nStarting DFS from vertex 2:")
	dfs(graph, 2)

	fmt.Println("\nStarting BFS from vertex 2:")
	bfs(graph, 2)
}
