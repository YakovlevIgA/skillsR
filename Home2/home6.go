package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)

	var chars []rune // срез возможных символов

	for i := 'A'; i <= 'Z'; i++ {
		chars = append(chars, i)
	}

	for i := 'a'; i <= 'z'; i++ {
		chars = append(chars, i)
	}

	for i := '0'; i <= '9'; i++ {
		chars = append(chars, i)
	}

	l := big.NewInt(int64(len(chars) - 1)) // Количество символов

	var pass string

	for i := 0; i < n; i++ {

		placeInChars, err := rand.Int(rand.Reader, l)
		if err != nil {
			fmt.Println("Ошибка при генерации случайного числа:", err)
			return
		}
		pass += string(chars[placeInChars.Int64()])
	}
	fmt.Println(pass)
}

// Проверить длинну сгенерированного пароля на соответствие ожидаемой.
// Проверить символы вывода на соответствие исходному пулу.
// Проверить, что пароли различаются между вызовами.
// Проверить на различных краевых вводах.
