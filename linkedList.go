package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

// Добавление элемента в начало списка
func (list *DoublyLinkedList) Prepend(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
	}
}

// Печать всех элементов списка с обоих направлений
func (list *DoublyLinkedList) PrintForward() {
	for node := list.Head; node != nil; node = node.Next {
		fmt.Printf("%d <-> ", node.Value)
	}
	fmt.Println("nil")
}

func (list *DoublyLinkedList) PrintBackward() {
	for node := list.Tail; node != nil; node = node.Prev {
		fmt.Printf("%d <-> ", node.Value)
	}
	fmt.Println("nil")
}

// Удаление элемента по значению
func (list *DoublyLinkedList) Delete(value int) bool {
	if list.Head == nil {
		return false // Список пуст
	}

	// Если нужно удалить первый элемент
	if list.Head.Value == value {
		if list.Head.Next != nil {
			list.Head = list.Head.Next
			list.Head.Prev = nil
		} else {
			list.Head = nil
			list.Tail = nil
		}
		return true
	}

	// Если нужно удалить последний элемент
	if list.Tail.Value == value {
		if list.Tail.Prev != nil {
			list.Tail = list.Tail.Prev
			list.Tail.Next = nil
		} else {
			list.Head = nil
			list.Tail = nil
		}
		return true
	}

	// Поиск элемента для удаления
	for node := list.Head; node != nil; node = node.Next {
		if node.Value == value {
			node.Prev.Next = node.Next
			if node.Next != nil {
				node.Next.Prev = node.Prev
			}
			return true
		}
	}

	return false // Элемент не найден
}

// Поиск элемента по значению
func (list *DoublyLinkedList) Search(value int) bool {
	for node := list.Head; node != nil; node = node.Next {
		if node.Value == value {
			return true
		}
	}
	return false
}

// Инвертирование списка
func (list *DoublyLinkedList) Reverse() {
	current := list.Head
	var prev *Node
	list.Tail = list.Head // Новый хвост — это старый головной элемент

	// Переворачиваем связи
	for current != nil {
		next := current.Next
		current.Next = prev
		current.Prev = next
		prev = current
		current = next
	}

	list.Head = prev // Новый головной элемент
}

func main() {
	list := &DoublyLinkedList{}
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)
	list.PrintForward()  // Вывод: 3 <-> 2 <-> 1 <-> nil
	list.PrintBackward() // Вывод: 1 <-> 2 <-> 3 <-> nil

	// Удаление элемента
	if list.Delete(2) {
		fmt.Println("Элемент 2 удален")
	} else {
		fmt.Println("Элемент 2 не найден")
	}
	list.PrintForward() // Вывод: 3 <-> 1 <-> nil

	// Поиск элемента
	if list.Search(1) {
		fmt.Println("Элемент 1 найден")
	} else {
		fmt.Println("Элемент 1 не найден")
	}

	// Инвертирование списка
	list.Reverse()
	list.PrintForward() // Вывод: 1 <-> 3 <-> nil
}
