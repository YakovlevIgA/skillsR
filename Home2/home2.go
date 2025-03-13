package main

import "fmt"

type MyError struct{}

func (MyError) Error() string {
	return "MyError!"
}
func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
func main() {
	var err *MyError  // указатель на nil
	errorHandler(err) // nil указатель не вызывает срабатывание метода, но является по своей сути интерфейсом в двумя полями: Тип и значение. Когда есть тип, условие != nil уже не выполняется
	err = &MyError{}  // создается экземпляр структуры
	errorHandler(err) // подтягивается метод, удовлетворяющий структуре
}
