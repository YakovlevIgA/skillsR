package main

import "fmt"

type Stack struct {
	elements []interface{} // Используем пустой интерфейс для универсальности
}

// Push добавляет элемент в стек
func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

// Pop удаляет верхний элемент из стека и возвращает его
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.elements) == 0 {
		return nil, false
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, true
}

func (s *Stack) Peek() string {
	if len(s.elements) == 0 {
		return "no values"
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	output := fmt.Sprintf("%v", element)

	return output
}

func main() {
	stack := &Stack{}
	stack.Push(1)
	stack.Push("Hello")
	stack.Push(3.14)
	stack.Push(3.14)

	// Выводим элементы разных типов
	fmt.Println(stack.Pop()) // Вывод: 3.14, true
	fmt.Println(stack.Pop()) // Вывод: Hello, true
	fmt.Println(stack.Peek())
}
