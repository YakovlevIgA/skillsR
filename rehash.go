package main

import (
	"fmt"
)

const initialSize = 7

type HashTable struct {
	array    []*bucket
	capacity int
	size     int
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

func NewHashTable(size int) *HashTable {
	table := &HashTable{
		array:    make([]*bucket, size),
		capacity: size,
		size:     0,
	}
	for i := range table.array {
		table.array[i] = &bucket{}
	}
	return table
}

func (h *HashTable) hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % h.capacity
}

func (h *HashTable) Insert(key string, value int) {
	index := h.hash(key)
	if !h.array[index].searchKey(key) {
		h.size++
	}
	h.array[index].insert(key, value)

	if float64(h.size)/float64(h.capacity) > 0.75 { // проверяем загруженность и делаем рехэш, если нужно
		h.rehash()
	}
}

func (h *HashTable) Search(key string) (int, bool) {
	index := h.hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := h.hash(key)
	if h.array[index].delete(key) {
		h.size--
	}
}

func (b *bucket) insert(key string, value int) {
	if !b.searchKey(key) {
		newNode := &bucketNode{key: key, value: value, next: b.head}
		b.head = newNode
	} else {
		currNode := b.head
		for currNode.key != key {
			currNode = currNode.next
		}
		currNode.value = value
	}
}

func (b *bucket) search(key string) (int, bool) {
	currNode := b.head
	for currNode != nil {
		if currNode.key == key {
			return currNode.value, true
		}
		currNode = currNode.next
	}
	return 0, false
}

func (b *bucket) delete(key string) bool {
	if b.head == nil {
		return false
	}
	if b.head.key == key {
		b.head = b.head.next
		return true
	}
	prevNode := b.head
	for prevNode.next != nil && prevNode.next.key != key {
		prevNode = prevNode.next
	}
	if prevNode.next == nil {
		return false
	}
	prevNode.next = prevNode.next.next
	return true
}

func (b *bucket) searchKey(key string) bool {
	currNode := b.head
	for currNode != nil {
		if currNode.key == key {
			return true
		}
		currNode = currNode.next
	}
	return false
}

func (h *HashTable) rehash() {
	newCapacity := h.capacity * 2
	newTable := NewHashTable(newCapacity)

	for _, b := range h.array {
		currNode := b.head
		for currNode != nil {
			newTable.Insert(currNode.key, currNode.value)
			currNode = currNode.next
		}
	}

	h.array = newTable.array
	h.capacity = newTable.capacity
	h.size = newTable.size

	fmt.Println("Рехеширование завершено! Новый размер таблицы:", h.capacity)
}

func (h *HashTable) Size() (int, int) {
	return h.size, h.capacity
}

func (h *HashTable) LoadFactor() float64 {
	return float64(h.size) / float64(h.capacity)
}

func main() {
	hashTable := NewHashTable(initialSize)

	hashTable.Insert("Alice", 100)
	hashTable.Insert("Bob", 200)
	hashTable.Insert("Charlie", 300)
	hashTable.Insert("David", 400)
	hashTable.Insert("Eve", 500)
	hashTable.Insert("Frank", 600)

	val, found := hashTable.Search("Alice")
	fmt.Println("Alice:", val, found)
	val, found = hashTable.Search("Bob")
	fmt.Println("Bob:", val, found)
	size, capacity := hashTable.Size()
	fmt.Println("Элементов в таблице:", size, "размер:", capacity)

	hashTable.Insert("Grace", 700)

	fmt.Println("После рехеширования:")
	val, found = hashTable.Search("Charlie")
	fmt.Println("Charlie:", val, found)
	val, found = hashTable.Search("David")
	fmt.Println("David:", val, found)
	size, capacity = hashTable.Size()
	fmt.Println("Элементов в таблице:", size, "размер:", capacity)
}
