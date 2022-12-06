package main

import "fmt"

/*
 a channel is a technique which allows to let one goroutine to send data to another goroutine. By default channel is bidirectional, means the goroutines can send or receive data through the same channel as shown in the below image:
*/

func main() {

	//using make create a channel or global variable
	//var ch chan int---nil channel
	ch := make(chan int)

	go func(a, b int) {
		defer close(ch)
		c := a + b

		//by using (chanel name) <- we can assign then value to channel
		(ch) <- c

	}(1, 2)

	//this is reciving data fron channel <-

	//result,bool :=chan
	val, ok := <-ch
	if ok {
		fmt.Println(val, ok)
	}

	//2
	// ch := make(chan int, 1)
	// go func() {
	// 	defer close(ch)
	// 	for i := 0; i < 5; i++ {
	// 		ch <- i
	// 	}
	// }()

	// //range

	// for v := range ch {
	// 	fmt.Printf("v: %v\n", v)
	// }

}
