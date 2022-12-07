package main

import (
	"fmt"
	"sync"
)

///execte at only one time

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	load := func() {
		fmt.Println("Run only onece")
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			once.Do(load)
			//	load()
		}()
	}
	wg.Wait()
}
