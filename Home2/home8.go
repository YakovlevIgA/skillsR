package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ch
		fmt.Println("Горутина прочитала из канала")
	}()

	fmt.Println("Главная горутина отправила данные в канал")
	ch <- true

	wg.Wait()

	close(ch)
}
