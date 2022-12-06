package main

import "fmt"

//ownership to avide the panic deadlock
func main() {

	Server := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		return ch

	}

	Clinet := func(ch <-chan int) {
		for v := range ch {
			fmt.Printf("This is %d\n", v)
		}
		fmt.Println("Done!...")
	}

	b := Server()
	Clinet(b)

}
