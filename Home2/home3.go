package main

import "fmt"

func main() {

	var foo []int
	var bar []int
	foo = append(foo, 1) // {} + 1 === {1} len, cap = 1, 1

	foo = append(foo, 2) // {1} + 2 === {1, 2} len, cap = 2, 2

	foo = append(foo, 3) // {1, 2} + 3 === {1, 2, 3} len, cap = 3, 4

	bar = append(foo, 4) // {1, 2, 3} + 4 === bar{1, 2, 3, 4} len, cap = 4, 4
	// здесь foo и bar указывает на один исходный массив, но
	// срез отображается в зависимости от своего len.
	foo = append(foo, 5) // {1, 2, 3} + 5 ===  len, cap = 4, 4
	// здесь не происходит создания нового массива, len(foo) увеличивается до 4,
	// захватывая при этом четвертую ячейку у bar и на ее место вставляется 5.

	fmt.Println(foo, bar)

}
