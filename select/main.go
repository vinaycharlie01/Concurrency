package main

import (
	"fmt"
	"sync"
)

func Person1(c chan string, wg *sync.WaitGroup) {
	//time.Sleep(1 * time.Second)
	defer wg.Done()
	c <- "Person1"
}

func Person2(c chan string, wg *sync.WaitGroup) {
	//stime.Sleep(1 * time.Second)
	defer wg.Done()
	c <- "Person2"
}

func main() {
	ch1 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(4)
	go Person1(ch1, &wg)
	go Person2(ch1, &wg)
	go Person1(ch1, &wg)
	go Person2(ch1, &wg)
	for i := 0; i < 4; i++ {
		select {
		case m := <-ch1:
			fmt.Println(m)
		case n := <-ch1:
			fmt.Println(n)

		}
	}
	wg.Wait()
	//fmt.Println(<-ch1)

}
