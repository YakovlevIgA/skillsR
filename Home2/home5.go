package main

import "fmt"

func main() {
	str := "Привет"
	runes := []rune(str)
	runes[2] = 'e'
	str = string(runes) // строки неизменяемы. Для преобразования можно сконвертировать в срез рун и обратно
	fmt.Println(str)
}
