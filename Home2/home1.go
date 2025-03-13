package main

import "fmt"

func printNumber(ptrToNumber interface{}) {
	if ptrToNumber == nil { // проверка на непосредственный nil
		fmt.Println("nil 1")
		return
	}

	if ptr, ok := ptrToNumber.(*int); ok {
		if ptr == nil { // проверка на указатель на nil
			fmt.Println("nil 2")
		} else {
			fmt.Println(*ptr)
		}
	} else {
		fmt.Println("invalid type")
	}
}

func main() {
	v := 10
	printNumber(&v)

	var pv *int
	printNumber(pv)

	pv = &v
	printNumber(pv)
}
