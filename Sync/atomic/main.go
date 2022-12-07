package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//atomic Lockless operations

func main() {

	var a int64
	var wg sync.WaitGroup
	//var mu sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			//mu.Lock()
			atomic.AddInt64(&a, 1)
			//mu.Unlock()
		}()

	}
	wg.Wait()
	fmt.Println(a)

}
