package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 3) // небуф. канал, сразу залочился, делаем буфер.
	wg := &sync.WaitGroup{}
	wg.Add(3) // сделали ожидание 3 горутин
	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done() // списали с дефера каждую горутину
			ch <- v * v
		}(i)
	}
	wg.Wait() // дождались
	close(ch) // закрыли канал
	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)
}
