package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	// Выбираем первый элемент как опорный
	pivotIndex := 0

	// Перемещаем опорный элемент в конец
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Разделяем на подмассивы меньше и больше опорного
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	// Возвращаем опорный элемент на место
	arr[left], arr[right] = arr[right], arr[left]

	// Рекурсивно сортируем подмассивы
	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	sample := []int{9, -3, 5, 1, 6, 8, -6, 1, 1}
	sorted := quickSort(sample)
	fmt.Println("Отсортированный массив:", sorted)
}
