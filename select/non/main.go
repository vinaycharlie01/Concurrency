package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "Message"
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("No message received")

		}

		fmt.Println("Processing...")
		time.Sleep(1500 * time.Millisecond)

	}
}
