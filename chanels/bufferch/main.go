package main

import "fmt"

func main() {
	//ch:=make(chan int capacity)
	// ch := make(chan []int, 1)
	// go func(a int) {
	// 	defer close(ch)
	// 	for i := 0; i < a; i++ {
	// 		b := []int{}
	// 		b = append(b, i)
	// 		ch <- b
	// 	}
	// 	for i := 0; i < a; i++ {
	// 		b := []int{}
	// 		b = append(b, i)
	// 		ch <- b
	// 	}
	// }(4)

	// for v := range ch {
	// 	fmt.Println(v)
	// }

	///2

	ch := make(chan int, 6)
	go func() {
		defer close(ch)
		for i := 0; i <= 6; i++ {
			ch <- i
		}
	}()

	//range

	for v := range ch {
		fmt.Printf("v: %v\n", v)
	}
}
