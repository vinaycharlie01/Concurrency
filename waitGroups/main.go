package main

import (
	"fmt"
	"sync"
)

func Data(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(10 + a)

}

func main() {
	var data int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		data++
	}()
	//wg.Wait()
	go Data(10, &wg)
	wg.Wait()
	if data == 1 {
		fmt.Println("The value is ", data)
	}

	//go Data(10, &wg)

}
