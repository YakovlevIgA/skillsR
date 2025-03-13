func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	mergedChannel := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			mergedChannel <- n
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(mergedChannel)
	}()
	return mergedChannel
}