package main

import (
	"fmt"
	"sync"
)

func main() {
	var Balance int
	var wg sync.WaitGroup
	var mu sync.Mutex
	deposit := func(ammount int) {
		mu.Lock()

		Balance += ammount
		mu.Unlock()
	}

	withdraw := func(ammount int) {
		mu.Lock()
		defer mu.Unlock()
		Balance -= ammount

	}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdraw(1)
		}()
	}
	wg.Wait()
	fmt.Println("total balance: ", Balance)
}
