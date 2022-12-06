package main

import (
	"fmt"
)

func Send(ch chan int) {

	//if not close then channeLIits showing  fatal error: all goroutines are asleep - deadlock!
	//defer close(ch)
	for i := 0; i < 5; i++ {
		ch <- i
	}

	// time.Sleep(100 * time.Millisecond)
}

func main() {
	c := make(chan int)
	c2 := make(chan int)

	go Send(c)
	go Send(c2)

	for v := range c {
		fmt.Println(v)
	}

	for v := range c2 {
		fmt.Println(v)
	}

}
