package main

import (
	"fmt"
)

// Graph структура, представляющая граф
type Graph struct {
	vertices []*Vertex
	directed bool // Если true, то граф направленный, если false - ненаправленный
}

// Vertex структура, представляющая вершину графа
type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

// AddVertex добавляет вершину в граф
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
		return
	}
	g.vertices = append(g.vertices, &Vertex{Key: k})
}

// AddEdge добавляет ребро между двумя вершинами в графе
func (g *Graph) AddEdge(from, to int) {
	// Получаем вершины
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// Проверяем, что вершины существуют
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v->%v)", from, to)
		fmt.Println(err.Error())
		return
	}
	// Проверяем, что ребро не существует
	if contains(fromVertex.Adjacent, to) {
		err := fmt.Errorf("existing edge (%v->%v)", from, to)
		fmt.Println(err.Error())
		return
	}
	// Добавляем ребро в зависимости от типа графа
	fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
	if !g.directed {
		toVertex.Adjacent = append(toVertex.Adjacent, fromVertex)
	}
}

// Удаление вершины
func (g *Graph) RemoveVertex(k int) {
	vertex := g.getVertex(k)
	if vertex == nil {
		fmt.Printf("Vertex %v not found\n", k)
		return
	}

	// Удаляем все рёбра, которые ссылаются на эту вершину
	for _, v := range g.vertices {
		for i, adj := range v.Adjacent {
			if adj.Key == k {
				v.Adjacent = append(v.Adjacent[:i], v.Adjacent[i+1:]...)
				break
			}
		}
	}

	// Удаляем саму вершину из графа
	for i, v := range g.vertices {
		if v.Key == k {
			g.vertices = append(g.vertices[:i], g.vertices[i+1:]...)
			break
		}
	}
}

// Удаление рёбер
func (g *Graph) RemoveEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Println("One or both vertices not found")
		return
	}

	// Удаляем ребро из fromVertex
	for i, adj := range fromVertex.Adjacent {
		if adj.Key == to {
			fromVertex.Adjacent = append(fromVertex.Adjacent[:i], fromVertex.Adjacent[i+1:]...)
			break
		}
	}

	// Если граф ненаправленный, удаляем также из toVertex
	if !g.directed {
		for i, adj := range toVertex.Adjacent {
			if adj.Key == from {
				toVertex.Adjacent = append(toVertex.Adjacent[:i], toVertex.Adjacent[i+1:]...)
				break
			}
		}
	}
}

// Вспомогательная функция для проверки существования вершины
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}

// Вспомогательная функция для получения вершины по ключу
func (g *Graph) getVertex(k int) *Vertex {
	for _, v := range g.vertices {
		if v.Key == k {
			return v
		}
	}
	return nil
}

// DFS (поиск в глубину)
func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
	g.dfsHelper(g.getVertex(start), visited)
}

// Вспомогательная функция DFS
func (g *Graph) dfsHelper(v *Vertex, visited map[int]bool) {
	if v == nil {
		return
	}
	if visited[v.Key] {
		return
	}
	visited[v.Key] = true
	fmt.Print(v.Key, " ")

	for _, adj := range v.Adjacent {
		g.dfsHelper(adj, visited)
	}
}

// BFS (поиск в ширину)
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []*Vertex{g.getVertex(start)}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if visited[vertex.Key] {
			continue
		}

		visited[vertex.Key] = true
		fmt.Print(vertex.Key, " ")

		for _, adj := range vertex.Adjacent {
			if !visited[adj.Key] {
				queue = append(queue, adj)
			}
		}
	}
}

func main() {
	// Пример использования
	graph := &Graph{directed: false} // false для ненаправленного графа, true для направленного
	for i := 1; i <= 5; i++ {
		graph.AddVertex(i)
	}
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(4, 2)
	graph.AddEdge(5, 2)
	graph.AddEdge(3, 5)

	// Выводим граф
	fmt.Println("Graph:")
	for _, v := range graph.vertices {
		fmt.Printf("Vertex %v : ", v.Key)
		for _, v := range v.Adjacent {
			fmt.Printf("%v ", v.Key)
		}
		fmt.Println()
	}

	// Поиск в глубину (DFS)
	fmt.Println("DFS from 1:")
	graph.DFS(1) // Ожидаемый вывод: 1 2 3 4 5
	fmt.Println()

	// Поиск в ширину (BFS)
	fmt.Println("BFS from 1:")
	graph.BFS(1) // Ожидаемый вывод: 1 2 3 4 5
	fmt.Println()

	// Удаление рёбер и вершин
	graph.RemoveEdge(2, 3)
	fmt.Println("Graph after removing edge (2->3):")
	for _, v := range graph.vertices {
		fmt.Printf("Vertex %v : ", v.Key)
		for _, v := range v.Adjacent {
			fmt.Printf("%v ", v.Key)
		}
		fmt.Println()
	}

	graph.RemoveVertex(5)
	fmt.Println("Graph after removing vertex 5:")
	for _, v := range graph.vertices {
		fmt.Printf("Vertex %v : ", v.Key)
		for _, v := range v.Adjacent {
			fmt.Printf("%v ", v.Key)
		}
		fmt.Println()
	}
}
