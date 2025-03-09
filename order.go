package main

import "fmt"

type Queue struct {
	elements []interface{}
}

// Enqueue добавляет элемент в конец очереди
func (q *Queue) Enqueue(value interface{}) {
	q.elements = append(q.elements, value)
}

// Dequeue удаляет элемент из начала очереди и возвращает его
func (q *Queue) Dequeue() (interface{}, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, true
}

func (q *Queue) Peek() string {
	if len(q.elements) == 0 {
		return "no values"
	}
	element := q.elements[0]

	output := fmt.Sprintf("%v", element)

	return output
}

func main() {
	queue := &Queue{}
	queue.Enqueue("string")
	queue.Enqueue(2)
	queue.Enqueue(3.15)
	fmt.Println(queue.Dequeue()) // Вывод: 1, true
	fmt.Println(queue.Dequeue()) // Вывод: 2, true
	fmt.Println(queue.Peek())
}
