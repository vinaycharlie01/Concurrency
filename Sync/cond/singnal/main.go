package main

import (
	"fmt"
	"sync"
)

//waiting for condition if false then display the result
var sharedrsc = make(map[string]interface{})

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedrsc) <= 0 {
			c.Wait()
		}
		fmt.Println(sharedrsc["rc1"])
		fmt.Println(sharedrsc["rc2"])
		c.L.Unlock()
	}()

	go func() {
		defer wg.Done()
		c.L.Lock()
		sharedrsc["rc1"] = "foo"
		c.Broadcast()
		c.L.Unlock()
	}()

	go func() {
		defer wg.Done()
		c.L.Lock()
		sharedrsc["rc2"] = "Hello"
		c.Broadcast()
		c.L.Unlock()
	}()
	wg.Wait()
}
