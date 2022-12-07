package main

import (
	"fmt"
	"sync"
)

var sharedrsc = make(map[string]interface{})

//condition variable waiting if condition if true waitng for false if false unlocked and result

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedrsc) < 1 {
			c.Wait()
		}
		fmt.Println(sharedrsc["rc1"])
		c.L.Unlock()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedrsc) < 2 {
			c.Wait()
		}
		fmt.Println(sharedrsc["rc2"])
		c.L.Unlock()
	}()
	c.L.Lock()
	sharedrsc["rc1"] = "foo"
	sharedrsc["rc2"] = "foo2"
	c.Broadcast()
	c.L.Unlock()

	wg.Wait()
}
