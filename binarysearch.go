package main

import (
	"fmt"
)

// Ordered - интерфейс для поддерживаемых типов
type Ordered interface {
	int | float64 | string
}

// BinarySearch выполняет бинарный поиск по отсортированному массиву
func BinarySearch[T Ordered](arr []T, target T) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	// Поиск в массиве чисел
	arrInt := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	targetInt := 23
	resultInt := BinarySearch(arrInt, targetInt)
	fmt.Printf("Целое число найдено на позиции: %d\n", resultInt)

	// Поиск в массиве строк
	arrStr := []string{"apple", "banana", "cherry", "date", "fig", "grape"}
	targetStr := "fig"
	resultStr := BinarySearch(arrStr, targetStr)
	fmt.Printf("Строка найдена на позиции: %d\n", resultStr)
}
