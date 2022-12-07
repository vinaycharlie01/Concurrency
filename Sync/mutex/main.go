package main

import (
	"fmt"
	"sync"
)

//protect the go routines data

func main() {
	//b := make(chan int)

	var a int
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			a = a + 1
			mu.Unlock()
		}()

	}
	wg.Wait()
	fmt.Println(a)
}
