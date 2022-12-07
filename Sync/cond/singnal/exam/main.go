package main

import (
	"fmt"
	"sync"
)

var sharedrsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedrsc) <= 0 {
			c.Wait()
		}
		fmt.Println(sharedrsc["rc1"])
		c.L.Unlock()
	}()
	c.L.Lock()
	sharedrsc["rc1"] = "foo"
	c.Signal()
	c.L.Unlock()

	wg.Wait()
}
