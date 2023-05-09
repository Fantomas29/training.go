package main

import (
	"fmt"
	"sync"
)

func readFromChannel(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	val := <-ch
	fmt.Printf("Received: %d\n", val)
}

func main() {
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go readFromChannel(ch, &wg)
	// go readFromChannel(ch, &wg)
	// go readFromChannel(ch, &wg)

	for i := 1; i <= 3; i++ {
		// time.Sleep(1 * time.Second)
		ch <- i
		fmt.Printf("Sent: %d\n", i)
		// time.Sleep(1 * time.Second)
	}

	close(ch)

	wg.Wait()
	fmt.Println("Done")
}
