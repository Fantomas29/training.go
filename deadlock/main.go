package main

import (
	"fmt"
	"time"
)

func hello(c chan string) {
	c <- "Hello there"
	fmt.Println("hello() finished")
}

func main() {
	c := make(chan string)
	go hello(c)
	s := <-c
	fmt.Printf("s =%v\n", s)

	c <- "Message from main"
	s = <-c
	fmt.Printf("s =%v\n", s)
	time.Sleep(1 * time.Second)
}
